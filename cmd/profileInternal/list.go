package profileInternal

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/towbee05/insighta/internal/app/profile"
	"github.com/towbee05/insighta/pkg/myTypes"
)

var (
	gender                string
	ageGroup              string
	countryId             string
	minAge                int
	maxAge                int
	minGenderProbability  float32
	minCountryProbability float32
	sortBy                string
	order                 string
	page                  int
	limit                 int
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		var filters = myTypes.Filters{
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
		return profile.ListResponse(filters)
	},
}

func init() {
	ListCmd.Flags().StringVar(&gender, "gender", "", "Filter by gender.")
	ListCmd.Flags().StringVar(&ageGroup, "age-group", "", "Filter by age groups ['child', 'teenager', 'adult', 'senior'].")
	ListCmd.Flags().StringVar(&countryId, "country", "", "Filter by country code.")
	ListCmd.Flags().IntVar(&minAge, "min-age", 0, "Filter by minimum age.")
	ListCmd.Flags().IntVar(&maxAge, "max-age", 0, "Filter by maximum age.")
	ListCmd.Flags().Float32Var(&minGenderProbability, "min-gender-probability", 0, "Filter by minimum gender probability.")
	ListCmd.Flags().Float32Var(&minCountryProbability, "min-country-probability", 0, "Filter by minimum country probability.")
	ListCmd.Flags().StringVar(&sortBy, "sort-by", "", "Sort by field ['age', 'country', 'name', 'gender probability', 'age group', 'created at' ].")
	ListCmd.Flags().StringVar(&order, "order", "asc", "order by ['desc', 'asc'].")
	ListCmd.Flags().IntVar(&page, "page", 1, "Go to page.")
	ListCmd.Flags().IntVar(&limit, "limit", 10, "Limit profiles from DB.")
}
