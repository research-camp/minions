package minion

import (
	"github.com/amirhnajafiz/minions/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	MinIO storage.Storage
}

func (h Handler) Download(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Upload(ctx *fiber.Ctx) error {
	return nil
}
