package router

import (
	"github.com/amirhnajafiz/minions/internal/config"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Cfg config.RouterConfig
}

func (h Handler) Get(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Put(ctx *fiber.Ctx) error {
	return nil
}
