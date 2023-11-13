package config

import "github.com/amirhnajafiz/minions/internal/storage"

func DefaultRouter() RouterConfig {
	return RouterConfig{
		Port:    80,
		Minions: []string{},
	}
}

func DefaultMinion() MinionConfig {
	return MinionConfig{
		Port: 80,
		MinIO: storage.Config{
			Endpoint:  "",
			AccessKey: "",
			SecretKey: "",
			Bucket:    "",
			SSL:       false,
		},
	}
}
