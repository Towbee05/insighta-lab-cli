package profileInternal

import (
	"github.com/spf13/cobra"
	"github.com/towbee05/insighta/internal/app/profile"
)

var SearchCmd = &cobra.Command{
	Use:     "search",
	Aliases: []string{"find"},
	Short:   "Search for parameters with string",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return profile.SearchProfile(args[0])
	},
}
