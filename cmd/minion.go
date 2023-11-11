package cmd

import (
	"github.com/amirhnajafiz/minions/internal/config"
	"github.com/amirhnajafiz/minions/internal/http/minion"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type Minion struct {
	Cfg config.Config
}

func (m Minion) Command() *cobra.Command {
	return &cobra.Command{
		Use: "minion",
		Run: func(cmd *cobra.Command, args []string) {
			m.main()
		},
	}
}

func (m Minion) main() {
	app := fiber.New()

	h := minion.Handler{}

	app.Get("/download", h.Download)
	app.Post("/upload", h.Upload)
}
