package cmd

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/minions/internal/config"
	"github.com/amirhnajafiz/minions/internal/http/router"
	"github.com/amirhnajafiz/minions/internal/metrics"

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

	// create metrics struct
	m := metrics.Metrics{}
	m.Init(len(cfg.Minions))

	// create new handler
	router.Handler{
		Cfg:     cfg,
		Metrics: &m,
	}.Register(app)

	// start the listener
	if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatal(err)
	}
}
