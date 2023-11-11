package cmd

import (
	"github.com/amirhnajafiz/minions/internal/http/minion"

	"github.com/gofiber/fiber/v2"
)

type Minion struct {
}

func (m Minion) main() {
	app := fiber.New()

	h := minion.Handler{}

	app.Get("/download", h.Download)
	app.Post("/upload", h.Upload)
}
