package cmd

import (
	"github.com/spf13/cobra"
	"github.com/towbee05/insighta/internal/app/auth"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout of current session",
	Run: func(cmd *cobra.Command, args []string) {
		auth.Logout()
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
