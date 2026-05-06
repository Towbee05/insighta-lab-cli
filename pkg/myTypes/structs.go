package myTypes

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type PolledResult struct {
	Status  string  `json:"status"`
	Message *string `json:"message,omitempty"`
	Data    *Token  `json:"data,omitempty"`
}

type GithubAuthData struct {
	Authentication_url string `json:"authentication_url"`
	State              string `json:"state"`
}

type Profile struct {
	Id                  uuid.UUID `json:"id"`
	Name                string    `json:"name"`
	Gender              string    `json:"gender"`
	Gender_probability  float32   `json:"gender_probability"`
	Age                 int32     `json:"age"`
	Age_group           string    `json:"age_group"`
	Country_id          string    `json:"country_id"`
	Country_name        string    `json:"country_name"`
	Country_probability float32   `json:"country_probability"`
	Created_at          time.Time `json:"created_at"`
}

type LinkInProfile struct {
	Self string `json:"self"`
	Next string `json:"next"`
	Prev string `json:"prev"`
}

type ProfileResponse struct {
	Status      string        `json:"status"`
	Page        int32         `json:"page"`
	Limit       int32         `json:"limit"`
	Total       int32         `json:"total"`
	Total_pages int32         `json:"total_pages"`
	Links       LinkInProfile `json:"links"`
	Data        []Profile     `json:"data"`
}

type ErrorResponse struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
}

type UnauthorizedResponse struct {
	Detail string `json:"detail"`
}

type RefreshTokenBody struct {
	RefreshToken string `json:"refresh_token"`
}

type Filters struct {
	Gender                string
	AgeGroup              string
	CountryID             string
	MinAge                int
	MaxAge                int
	MinGenderProbability  float32
	MinCountryProbability float32
	SortBy                string
	Order                 string
	Page                  int
	Limit                 int
}

type ExportFilters struct {
	Format                string
	Gender                string
	AgeGroup              string
	CountryID             string
	MinAge                int
	MaxAge                int
	MinGenderProbability  float32
	MinCountryProbability float32
	SortBy                string
	Order                 string
	Page                  int
	Limit                 int
}

type PersonResponse struct {
	Status  string  `json:"status"`
	Message *string `json:"message"`
	Data    Profile `json:"data"`
}

type CreatePerson struct {
	Name string `json:"name"`
}
