package models

type Host struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	IP       string   `json:"ip"`
	SSHPort  int      `json:"sshPort"`
	Username string   `json:"username"`
	AuthType string   `json:"authType"`
	Tags     []string `json:"tags"`
}

type Pack struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Image        string   `json:"image"`
	InternalPort int      `json:"internalPort"`
	EnvSchema    []string `json:"envSchema"`
	Provides     []string `json:"provides"`
	Requires     []string `json:"requires"`
}

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Environment string `json:"environment"`
	Description string `json:"description"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type GraphNode struct {
	ID           string         `json:"id"`
	Type         string         `json:"type"`
	PackID       string         `json:"packId"`
	Label        string         `json:"label"`
	HostID       string         `json:"hostId"`
	InternalPort int            `json:"internalPort"`
	Position     Position       `json:"position"`
	Config       map[string]any `json:"config"`
}

type GraphEdge struct {
	ID     string `json:"id"`
	Source string `json:"source"`
	Target string `json:"target"`
	Type   string `json:"type"`
}

type ProjectGraph struct {
	ProjectID string      `json:"projectId"`
	Nodes     []GraphNode `json:"nodes"`
	Edges     []GraphEdge `json:"edges"`
}

type Exposure struct {
	NodeID       string `json:"nodeId"`
	Enabled      bool   `json:"enabled"`
	ExternalPort *int   `json:"externalPort"`
	InternalPort int    `json:"internalPort"`
}

type InstallConfig struct {
	ProjectID string     `json:"projectId"`
	Exposures []Exposure `json:"exposures"`
}

type GeneratedOutput struct {
	ProjectID          string            `json:"projectId"`
	Environment        map[string]string `json:"environment"`
	NginxConfigPreview string            `json:"nginxConfigPreview"`
	DeployCommands     []string          `json:"deployCommands"`
}
