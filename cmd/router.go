package cmd

import (
	"github.com/amirhnajafiz/minions/internal/config"

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

}
