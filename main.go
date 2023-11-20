package main

import (
	"log"

	"github.com/amirhnajafiz/minions/cmd"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {
	// load env file if existed
	_ = godotenv.Load()

	// create the root command
	root := &cobra.Command{}

	// add sub commands
	root.AddCommand(
		cmd.Minion{}.Command(),
		cmd.Router{}.Command(),
	)

	// execute root command
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
