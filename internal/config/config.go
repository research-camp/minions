package config

import "github.com/amirhnajafiz/minions/internal/storage"

type (
	MinionConfig struct {
		Port   int
		Router string
		MinIO  storage.Config
	}

	RouterConfig struct {
		Port    int
		Minions []string
	}
)

func LoadMinion() MinionConfig {
	return DefaultMinion()
}

func LoadRouter() RouterConfig {
	return DefaultRouter()
}
