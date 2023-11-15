package router

import (
	"github.com/amirhnajafiz/minions/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Handler struct {
	Cfg     config.RouterConfig
	Metrics Metrics
}

func (h Handler) Register(app *fiber.App) {
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	app.Get("/", h.Signal)

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Get("/get", h.Get)
	app.Post("/put", h.Put)
}
