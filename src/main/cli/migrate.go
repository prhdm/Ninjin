package cli

import (
	"github.com/spf13/cobra"
	"farm/src/models"
)

var migrate = &cobra.Command{
	Use:  "migrate",
	Run: func(cobraCmd *cobra.Command, args []string) {
		models.PostgresDBProvider.MigrateDB()
	},
}

func init() {
	rootCmd.AddCommand(migrate)
}
