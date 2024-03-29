package cmd

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/minions/internal/config"
	"github.com/amirhnajafiz/minions/internal/http/minion"
	"github.com/amirhnajafiz/minions/internal/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
)

type Minion struct{}

func (m Minion) Command() *cobra.Command {
	return &cobra.Command{
		Use: "minion",
		Run: func(cmd *cobra.Command, args []string) {
			m.main()
		},
	}
}

func (m Minion) main() {
	// load configs
	cfg := config.LoadMinion()

	// create new fiber
	app := fiber.New()

	// open connection to MinIO
	client, err := storage.New(cfg.MinIO)
	if err != nil {
		log.Fatal(err)
	}

	// create a new handler
	minion.Handler{
		MinIO: client,
	}.Register(app)

	// start the listener
	if er := app.Listen(fmt.Sprintf(":%d", cfg.Port)); er != nil {
		log.Fatal(er)
	}
}
