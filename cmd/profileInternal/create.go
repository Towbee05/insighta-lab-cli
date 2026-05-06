package profileInternal

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/towbee05/insighta/internal/app/profile"
)

var name string

var CreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"make", "post"},
	Short:   "Create a user with provided name",
	RunE: func(cmd *cobra.Command, args []string) error {
		if name == "" {
			return fmt.Errorf("Please provide a name")
		}
		return profile.CreateProfile(name)
	},
}

func init() {
	CreateCmd.Flags().StringVar(&name, "name", "", "Add profile to DB")
}
