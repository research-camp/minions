package config

import "github.com/amirhnajafiz/minions/internal/storage"

type (
	MinionConfig struct {
		Port  int            `koanf:"port"`
		MinIO storage.Config `koanf:"minio"`
	}

	RouterConfig struct {
		Port    int      `koanf:"port"`
		Minions []string `koanf:"minions"`
	}
)
