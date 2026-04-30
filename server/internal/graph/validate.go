package graph

import "om-platform/server/internal/models"

var allowedEdges = map[string]map[string]bool{
	"nginx":       {"frontend": true, "backend": true},
	"frontend":    {"backend": true},
	"backend":     {"db-postgres": true, "db-mysql": true, "redis": true},
	"db-postgres": {},
	"db-mysql":    {},
	"redis":       {},
}

func ValidateTopology(g models.ProjectGraph) []string {
	errors := []string{}
	nodeByID := map[string]models.GraphNode{}
	for _, n := range g.Nodes {
		nodeByID[n.ID] = n
	}
	for _, e := range g.Edges {
		src, sok := nodeByID[e.Source]
		tgt, tok := nodeByID[e.Target]
		if !sok || !tok {
			errors = append(errors, "edge references missing node: "+e.ID)
			continue
		}
		if !allowedEdges[src.Type][tgt.Type] {
			errors = append(errors, "invalid edge: "+src.Type+" -> "+tgt.Type)
		}
	}
	return errors
}
