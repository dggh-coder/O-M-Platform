package storage

import (
	"fmt"
	"sync"

	"om-platform/server/internal/models"
)

type MemoryStore struct {
	mu       sync.RWMutex
	hosts    map[string]models.Host
	packs    map[string]models.Pack
	projects map[string]models.Project
	graphs   map[string]models.ProjectGraph
	installs map[string]models.InstallConfig
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		hosts:    map[string]models.Host{},
		packs:    map[string]models.Pack{},
		projects: map[string]models.Project{},
		graphs:   map[string]models.ProjectGraph{},
		installs: map[string]models.InstallConfig{},
	}
}

func (s *MemoryStore) UpsertHost(h models.Host) { s.mu.Lock(); defer s.mu.Unlock(); s.hosts[h.ID] = h }
func (s *MemoryStore) ListHosts() []models.Host {
	s.mu.RLock()
	defer s.mu.RUnlock()
	res := make([]models.Host, 0, len(s.hosts))
	for _, h := range s.hosts {
		res = append(res, h)
	}
	return res
}

func (s *MemoryStore) UpsertPack(p models.Pack) { s.mu.Lock(); defer s.mu.Unlock(); s.packs[p.ID] = p }
func (s *MemoryStore) ListPacks() []models.Pack {
	s.mu.RLock()
	defer s.mu.RUnlock()
	res := make([]models.Pack, 0, len(s.packs))
	for _, p := range s.packs {
		res = append(res, p)
	}
	return res
}

func (s *MemoryStore) UpsertProject(p models.Project) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.projects[p.ID] = p
}
func (s *MemoryStore) ListProjects() []models.Project {
	s.mu.RLock()
	defer s.mu.RUnlock()
	res := make([]models.Project, 0, len(s.projects))
	for _, p := range s.projects {
		res = append(res, p)
	}
	return res
}

func (s *MemoryStore) SetGraph(projectID string, g models.ProjectGraph) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.graphs[projectID] = g
}
func (s *MemoryStore) GetGraph(projectID string) (models.ProjectGraph, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	g, ok := s.graphs[projectID]
	if !ok {
		return models.ProjectGraph{}, fmt.Errorf("graph not found")
	}
	return g, nil
}

func (s *MemoryStore) SetInstall(projectID string, c models.InstallConfig) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.installs[projectID] = c
}
func (s *MemoryStore) GetInstall(projectID string) (models.InstallConfig, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	c, ok := s.installs[projectID]
	if !ok {
		return models.InstallConfig{}, fmt.Errorf("install config not found")
	}
	return c, nil
}
