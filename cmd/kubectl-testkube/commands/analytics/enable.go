package analytics

import (
	"github.com/kubeshop/testkube/cmd/kubectl-testkube/config"
	"github.com/kubeshop/testkube/pkg/ui"
	"github.com/spf13/cobra"
)

func NewEnableAnalyticsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "enable",
		Aliases: []string{"off", "e", "y"},
		Short:   "Enable collecting of anonymous analytics",
		Run: func(cmd *cobra.Command, args []string) {
			ui.Logo()
			config.Config.EnableAnalytics()
			err := config.Config.Save(config.Config.Data)
			ui.ExitOnError("saving config file", err)
			ui.Success("Analytics", "enabled")
		},
	}

	return cmd
}