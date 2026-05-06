package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "insighta",
	Short: "This is insighta lab+",
	Long:  "This is insighta lab+. Login via github to enjoy what we have to offer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Opps, an error occured while trying to execute insighta lab '%s'\n", err)
		os.Exit(1)
	}
}
