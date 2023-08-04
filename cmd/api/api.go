package api

import (
	"github.com/kyrinkie/admgo/src/api"
	"github.com/spf13/cobra"
)

var (
	configYml string
	apiCheck  bool
	StartCmd  = &cobra.Command{
		Use:   "api",
		Short: "Start API server",
		// Example:      "admgo api -c config/settings.yml",
		SilenceUsage: true,
		// PreRun: func(cmd *cobra.Command, args []string) {
		// 	setup()
		// },
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	api.Run()
	return nil
}
