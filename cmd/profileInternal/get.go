package profileInternal

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/towbee05/insighta/internal/app/profile"
)

var GetCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"fetch"},
	Short:   "Get [id].",
	Long:    "Get user by ID.",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := uuid.Parse(args[0])
		if err != nil {
			return fmt.Errorf("failed to convert %s to uuid ", args[0])
		}
		return profile.GetProfile(id)
	},
}
