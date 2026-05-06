package utils

import (
	"net/url"
	"strconv"

	"github.com/towbee05/insighta/pkg/myTypes"
)

func BuildQueryParams(filters myTypes.Filters) string {
	params := url.Values{}

	if filters.AgeGroup != "" {
		params.Add("age_group", filters.AgeGroup)
	}
	if filters.CountryID != "" {
		params.Add("country_id", filters.CountryID)
	}
	if filters.Gender != "" {
		params.Add("gender", filters.Gender)
	}
	if filters.Limit > 0 {
		params.Add("limit", strconv.Itoa(filters.Limit))
	}
	if filters.MaxAge > 0 {
		params.Add("max_age", strconv.Itoa(filters.MaxAge))
	}
	if filters.MinAge > 0 {
		params.Add("min_age", strconv.Itoa(filters.MinAge))
	}
	if filters.MinCountryProbability > 0 {
		params.Add("min_country_probability", strconv.FormatFloat(float64(filters.MinCountryProbability), 'f', -1, 32))
	}
	if filters.MinGenderProbability > 0 {
		params.Add("min_gender_probability", strconv.FormatFloat(float64(filters.MinGenderProbability), 'f', -1, 32))
	}
	if filters.Order != "" {
		params.Add("order", filters.Order)
	}
	if filters.Page > 0 {
		params.Add("page", strconv.Itoa(filters.Page))
	}
	if filters.SortBy != "" {
		params.Add("sort_by", filters.SortBy)
	}

	return params.Encode()
}
