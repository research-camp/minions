package minion

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const LocalDir = "./tmp/local"

type Handler struct {
	Router string
	MinIO  MinIO
}

func (h Handler) Register(app *fiber.App) {
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/download", h.download)
	app.Post("/upload", h.upload)
}
