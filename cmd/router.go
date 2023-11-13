package cmd

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/minions/internal/config"
	"github.com/amirhnajafiz/minions/internal/http/router"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type Router struct{}

func (r Router) Command() *cobra.Command {
	return &cobra.Command{
		Use: "router",
		Run: func(cmd *cobra.Command, args []string) {
			r.main()
		},
	}
}

func (r Router) main() {
	// load configs
	cfg := config.LoadRouter()

	// create new fiber
	app := fiber.New()

	// create new handler
	h := router.Handler{
		Cfg: cfg,
	}

	app.Get("/get", h.Get)
	app.Post("/put", h.Put)

	// start the listener
	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatal(err)
	}
}
