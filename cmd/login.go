package cmd

import (
	"github.com/spf13/cobra"
	"github.com/towbee05/insighta/internal/app/auth"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the application",
	Long:  "Login to the application using your GitHub account",
	Run: func(cmd *cobra.Command, args []string) {
		auth.Login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
