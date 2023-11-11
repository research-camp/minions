package minion

import "github.com/amirhnajafiz/minions/internal/storage"

type Handler struct {
	MinIO storage.Storage
}
