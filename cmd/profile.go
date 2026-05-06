package cmd

import (
	"github.com/spf13/cobra"
	profileInternal "github.com/towbee05/insighta/cmd/profileInternal"
)

var profileCmd = &cobra.Command{
	Use:     "profile",
	Aliases: []string{"profiles"},
	Short:   "Handle profiles from db",
}

func init() {
	rootCmd.AddCommand(profileCmd)
	profileCmd.AddCommand(profileInternal.GetCmd)
	profileCmd.AddCommand(profileInternal.ListCmd)
	profileCmd.AddCommand(profileInternal.SearchCmd)
	profileCmd.AddCommand(profileInternal.CreateCmd)
	profileCmd.AddCommand(profileInternal.ExportCmd)
}
