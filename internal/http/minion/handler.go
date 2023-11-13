package minion

import (
	"fmt"

	"github.com/amirhnajafiz/minions/internal/storage"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	MinIO *storage.Storage
}

func (h Handler) Download(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Upload(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return fmt.Errorf("failed to get multipart form: %w", err)
	}

	for _, file := range form.File["file"] {
		if er := ctx.SaveFile(file, fmt.Sprintf("./%s", file.Filename)); er != nil {
			return fmt.Errorf("failed to save file on local: %w", er)
		}

		if er := h.MinIO.Put(file.Filename, file.Filename); er != nil {
			return fmt.Errorf("failed to send on MinIo: %w", er)
		}
	}

	return ctx.SendStatus(fiber.StatusOK)
}
