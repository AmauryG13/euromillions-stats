package config

type Log struct {
	Name  string
	Level string
	File  string
}

type Service struct {
	Name string
}

type GRPC struct {
	Namespace string
	Address   string
}
