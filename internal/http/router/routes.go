package router

import (
	"fmt"

	"github.com/amirhnajafiz/minions/pkg/enum"

	"github.com/gofiber/fiber/v2"
)

func (h Handler) metricsHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(h.Metrics.Pull())
}

func (h Handler) signal(ctx *fiber.Ctx) error {
	key := ctx.Query("signal")

	switch key {
	case enum.HitSignal:
		h.Metrics.Hit()
	case enum.MissSignal:
		h.Metrics.Miss()
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (h Handler) get(ctx *fiber.Ctx) error {
	name := ctx.Query("file", "")
	if len(name) == 0 {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	if len(h.Cfg.Minions) == 0 {
		return fmt.Errorf("failed to get any cache")
	}

	index := len(name) % len(h.Cfg.Minions)
	url := h.Cfg.Minions[index]

	h.Metrics.Down()

	return ctx.Redirect(fmt.Sprintf("%s/download", url))
}

func (h Handler) put(ctx *fiber.Ctx) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return fmt.Errorf("failed to get multipart form: %w", err)
	}

	url := ""

	for _, file := range form.File["file"] {
		index := len(file.Filename) % len(h.Cfg.Minions)

		url = h.Cfg.Minions[index]
	}

	if len(url) == 0 {
		return fmt.Errorf("failed to get any cache: %w", err)
	}

	h.Metrics.Up()

	return ctx.Redirect(fmt.Sprintf("%s/upload", url))
}
