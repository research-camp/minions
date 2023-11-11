package main

import (
	"github.com/amirhnajafiz/minions/cmd"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{}

	root.AddCommand(
		cmd.Minion{}.Command(),
		cmd.Router{}.Command(),
	)

	if err := root.Execute(); err != nil {
		panic(err)
	}
}
