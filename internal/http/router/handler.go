package router

import "github.com/gofiber/fiber/v2"

type Handler struct{}

func (h Handler) Get(ctx *fiber.Ctx) error {
	return nil
}

func (h Handler) Put(ctx *fiber.Ctx) error {
	return nil
}
