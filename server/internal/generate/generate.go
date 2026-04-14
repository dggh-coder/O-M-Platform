package generate

import (
	"fmt"

	"om-platform/server/internal/models"
)

func FromGraph(g models.ProjectGraph, install models.InstallConfig) models.GeneratedOutput {
	env := map[string]string{}
	commands := []string{}
	for _, edge := range g.Edges {
		src := findNode(g.Nodes, edge.Source)
		tgt := findNode(g.Nodes, edge.Target)
		if src == nil || tgt == nil {
			continue
		}
		switch {
		case src.Type == "frontend" && tgt.Type == "backend":
			env["API_BASE_URL"] = fmt.Sprintf("http://%s:%d", tgt.ID, tgt.InternalPort)
		case src.Type == "backend" && (tgt.Type == "db-postgres" || tgt.Type == "db-mysql"):
			env["DB_HOST"] = tgt.ID
			env["DB_PORT"] = fmt.Sprintf("%d", tgt.InternalPort)
		case src.Type == "backend" && tgt.Type == "redis":
			env["REDIS_HOST"] = tgt.ID
			env["REDIS_PORT"] = fmt.Sprintf("%d", tgt.InternalPort)
		}
	}

	for _, n := range g.Nodes {
		commands = append(commands, fmt.Sprintf("docker run -d --name %s %s", n.ID, n.PackID))
	}

	_ = install
	return models.GeneratedOutput{
		ProjectID:          g.ProjectID,
		Environment:        env,
		NginxConfigPreview: "upstream frontend_upstream { server frontend_1:8080; }",
		DeployCommands:     commands,
	}
}

func findNode(nodes []models.GraphNode, id string) *models.GraphNode {
	for i := range nodes {
		if nodes[i].ID == id {
			return &nodes[i]
		}
	}
	return nil
}
