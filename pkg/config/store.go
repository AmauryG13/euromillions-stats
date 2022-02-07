package config

type StoreNode struct {
	Username string
	Password string
	Port     string
}

type Store struct {
	Nodes    []StoreNode
	Database string
	Table    string
}
