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

	cfg.Port, _ = strconv.Atoi(os.Getenv("MI_PORT"))
	cfg.Router = os.Getenv("MI_ROUTER")
	cfg.MinIO.Endpoint = os.Getenv("MI_MINIO_ENDPOINT")
	cfg.MinIO.AccessKey = os.Getenv("MI_MINIO_ACCESS")
	cfg.MinIO.SecretKey = os.Getenv("MI_MINIO_SECRET")
	cfg.MinIO.Bucket = os.Getenv("MI_MINIO_BUCKET")

	if os.Getenv("MI_MINIO_SSL") == "true" {
		cfg.MinIO.SSL = true
	}

	return cfg
}

func LoadRouter() RouterConfig {
	cfg := DefaultRouter()

	cfg.Port, _ = strconv.Atoi(os.Getenv("RT_PORT"))
	cfg.Minions = strings.Split(os.Getenv("RT_MINIONS"), ",")

	return cfg
}
