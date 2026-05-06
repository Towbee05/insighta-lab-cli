package profileInternal

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/towbee05/insighta/internal/app/profile"
	"github.com/towbee05/insighta/pkg/myTypes"
)

var format string

var ExportCmd = &cobra.Command{
	Use:     "export",
	Aliases: []string{"download"},
	Short:   "Export data from DB into PC",
	RunE: func(cmd *cobra.Command, args []string) error {
		var filters = myTypes.ExportFilters{
			Format:                format,
			Gender:                gender,
			AgeGroup:              ageGroup,
			CountryID:             countryId,
			MinAge:                minAge,
			MaxAge:                maxAge,
			MinGenderProbability:  minGenderProbability,
			MinCountryProbability: minCountryProbability,
			SortBy:                sortBy,
			Order:                 order,
			Page:                  page,
			Limit:                 limit,
		}
		if order != "asc" && order != "desc" {
			return fmt.Errorf("Order must be asc or desc")
		}
		return profile.ExportProfile(filters)
	},
}

func init() {
	ExportCmd.Flags().StringVar(&format, "format", "csv", "Add format to export ['csv','json','pdf']. ")
	ExportCmd.Flags().StringVar(&gender, "gender", "", "Filter by gender.")
	ExportCmd.Flags().StringVar(&ageGroup, "age-group", "", "Filter by age groups ['child', 'teenager', 'adult', 'senior'].")
	ExportCmd.Flags().StringVar(&countryId, "country", "", "Filter by country code.")
	ExportCmd.Flags().IntVar(&minAge, "min-age", 0, "Filter by minimum age.")
	ExportCmd.Flags().IntVar(&maxAge, "max-age", 0, "Filter by maximum age.")
	ExportCmd.Flags().Float32Var(&minGenderProbability, "min-gender-probability", 0, "Filter by minimum gender probability.")
	ExportCmd.Flags().Float32Var(&minCountryProbability, "min-country-probability", 0, "Filter by minimum country probability.")
	ExportCmd.Flags().StringVar(&sortBy, "sort-by", "", "Sort by field ['age', 'country', 'name', 'gender probability', 'age group', 'created at' ].")
	ExportCmd.Flags().StringVar(&order, "order", "asc", "order by ['desc', 'asc'].")
	ExportCmd.Flags().IntVar(&page, "page", 1, "Go to page.")
	ExportCmd.Flags().IntVar(&limit, "limit", 10, "Limit profiles from DB.")
}
