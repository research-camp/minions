package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/amirhnajafiz/minions/internal/storage"
)

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
	cfg := DefaultMinion()

	return cfg
}

func LoadRouter() RouterConfig {
	cfg := DefaultRouter()

	cfg.Port, _ = strconv.Atoi(os.Getenv("MI_PORT"))
	cfg.Minions = strings.Split(os.Getenv("MI_MINIONS"), ",")

	return cfg
}
