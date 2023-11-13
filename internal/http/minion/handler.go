package minion

import (
	"fmt"
	"os"

	"github.com/amirhnajafiz/minions/internal/storage"

	"github.com/gofiber/fiber/v2"
)

const LocalDir = "./tmp/local"

type Handler struct {
	MinIO *storage.Storage
}

func (h Handler) Download(ctx *fiber.Ctx) error {
	name := ctx.Query("file", "")
	if len(name) == 0 {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	code := fiber.StatusOK
	path := fmt.Sprintf("%s/%s", LocalDir, name)

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if er := h.MinIO.Get(name, path); er != nil {
				return ctx.SendStatus(fiber.StatusInternalServerError)
			}

			code = fiber.StatusCreated
		} else {
			return ctx.SendStatus(fiber.StatusNotFound)
		}
	}

	return ctx.Status(code).SendFile(path)
}

func (h Handler) Upload(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return fmt.Errorf("failed to get multipart form: %w", err)
	}

	for _, file := range form.File["file"] {
		path := fmt.Sprintf("%s/%s", LocalDir, file.Filename)
		if er := ctx.SaveFile(file, path); er != nil {
			return fmt.Errorf("failed to save file on local: %w", er)
		}

		if er := h.MinIO.Put(file.Filename, path); er != nil {
			return fmt.Errorf("failed to send on MinIo: %w", er)
		}
	}

	return ctx.SendStatus(fiber.StatusOK)
}
