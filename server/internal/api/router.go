package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"om-platform/server/internal/generate"
	"om-platform/server/internal/graph"
	"om-platform/server/internal/models"
	"om-platform/server/internal/storage"
)

func NewRouter() http.Handler {
	store := storage.NewMemoryStore()
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("/api/hosts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, store.ListHosts())
		case http.MethodPost:
			var payload models.Host
			if !decode(w, r, &payload) {
				return
			}
			store.UpsertHost(payload)
			writeJSON(w, http.StatusCreated, payload)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/packs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, store.ListPacks())
		case http.MethodPost:
			var payload models.Pack
			if !decode(w, r, &payload) {
				return
			}
			store.UpsertPack(payload)
			writeJSON(w, http.StatusCreated, payload)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/projects", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			writeJSON(w, http.StatusOK, store.ListProjects())
		case http.MethodPost:
			var payload models.Project
			if !decode(w, r, &payload) {
				return
			}
			store.UpsertProject(payload)
			writeJSON(w, http.StatusCreated, payload)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/api/projects/", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/projects/")
		parts := strings.Split(path, "/")
		if len(parts) < 2 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		projectID := parts[0]
		resource := parts[1]
		suffix := ""
		if len(parts) > 2 {
			suffix = parts[2]
		}

		switch resource {
		case "graph":
			if suffix == "validate" {
				handleGraphValidate(w, r, store, projectID)
				return
			}
			if suffix == "generate" {
				handleGenerate(w, r, store, projectID)
				return
			}
			handleGraph(w, r, store, projectID)
		case "install-config":
			if suffix == "validate" {
				handleInstallValidate(w, r, store, projectID)
				return
			}
			handleInstallConfig(w, r, store, projectID)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	return mux
}

func handleGraph(w http.ResponseWriter, r *http.Request, store *storage.MemoryStore, projectID string) {
	switch r.Method {
	case http.MethodGet:
		g, err := store.GetGraph(projectID)
		if err != nil {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, g)
	case http.MethodPut:
		var payload models.ProjectGraph
		if !decode(w, r, &payload) {
			return
		}
		payload.ProjectID = projectID
		store.SetGraph(projectID, payload)
		writeJSON(w, http.StatusOK, payload)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleGraphValidate(w http.ResponseWriter, r *http.Request, store *storage.MemoryStore, projectID string) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	g, err := store.GetGraph(projectID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	errs := graph.ValidateTopology(g)
	writeJSON(w, http.StatusOK, map[string]any{"valid": len(errs) == 0, "errors": errs})
}

func handleGenerate(w http.ResponseWriter, r *http.Request, store *storage.MemoryStore, projectID string) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	g, gErr := store.GetGraph(projectID)
	install, iErr := store.GetInstall(projectID)
	if gErr != nil || iErr != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "graph/install-config missing"})
		return
	}
	writeJSON(w, http.StatusOK, generate.FromGraph(g, install))
}

func handleInstallConfig(w http.ResponseWriter, r *http.Request, store *storage.MemoryStore, projectID string) {
	switch r.Method {
	case http.MethodGet:
		cfg, err := store.GetInstall(projectID)
		if err != nil {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusOK, cfg)
	case http.MethodPut:
		var payload models.InstallConfig
		if !decode(w, r, &payload) {
			return
		}
		payload.ProjectID = projectID
		store.SetInstall(projectID, payload)
		writeJSON(w, http.StatusOK, payload)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleInstallValidate(w http.ResponseWriter, r *http.Request, store *storage.MemoryStore, projectID string) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	cfg, err := store.GetInstall(projectID)
	if err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	used := map[int]string{}
	errs := []string{}
	for _, ex := range cfg.Exposures {
		if !ex.Enabled || ex.ExternalPort == nil {
			continue
		}
		if owner, ok := used[*ex.ExternalPort]; ok {
			errs = append(errs, "port collision: "+owner+" and "+ex.NodeID)
			continue
		}
		used[*ex.ExternalPort] = ex.NodeID
	}
	writeJSON(w, http.StatusOK, map[string]any{"valid": len(errs) == 0, "errors": errs})
}

func decode(w http.ResponseWriter, r *http.Request, dst any) bool {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return false
	}
	return true
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
