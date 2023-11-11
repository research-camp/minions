package cmd

import (
	"log"

	"github.com/amirhnajafiz/minions/internal/config"
	"github.com/amirhnajafiz/minions/internal/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type Router struct {
	Cfg config.Config
}

func (r Router) Command() *cobra.Command {
	return &cobra.Command{
		Use: "router",
		Run: func(cmd *cobra.Command, args []string) {
			r.main()
		},
	}
}

func (r Router) main() {
	app := fiber.New()

	h := router.Handler{}

	app.Get("/get", h.Get)
	app.Post("/put", h.Put)

	if err := app.Listen(":80"); err != nil {
		log.Fatal(err)
	}
}
