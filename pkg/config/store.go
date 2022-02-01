package config

type NodeOptions struct {
	Username string
	Password string
	Port     string
}

type Store struct {
	Nodes    []NodeOptions
	Database string
	Table    string
}
