package router

import (
	"github.com/amirhnajafiz/minions/internal/config"
	"github.com/amirhnajafiz/minions/internal/metrics"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Cfg     config.RouterConfig
	Metrics *metrics.Metrics
}

func (h Handler) Get(ctx *fiber.Ctx) error {
	h.Metrics.Down()

	return nil
}

func (h Handler) Put(ctx *fiber.Ctx) error {
	h.Metrics.Up()

	return nil
}
