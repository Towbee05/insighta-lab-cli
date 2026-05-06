package utils

import (
	"fmt"
	"net/http"

)

func AddTokenToRequestHeader(request *http.Request) error {
	token, tokenErr := CheckTokenFileExistence()
	if tokenErr != nil {
		return fmt.Errorf("an error occured generating token: %s \n", tokenErr)
	}

	request, requestErr := http.NewRequest("GET", "http://localhost:8000/api/profiles?limit=15&page=3", nil)
	if requestErr != nil {
		return fmt.Errorf("An error occured sending request: %s \n", requestErr)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	request.Header.Add("X-API-Version", "1")
	request.Header.Add("Content-Type", "application/json")

	return nil
}
