package profile

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/towbee05/insighta/pkg/myTypes"
	"github.com/towbee05/insighta/pkg/utils"
)

func GetProfile(id uuid.UUID) error {
	fmt.Println(id)
	token, tokenErr := utils.CheckTokenFileExistence()
	if tokenErr != nil {
		return fmt.Errorf("an error occured generating token: %s \n", tokenErr)
	}
	var url string = fmt.Sprintf("http://localhost:8000/api/profiles/%s", id)

	request, requestErr := http.NewRequest("GET", url, nil)
	if requestErr != nil {
		return fmt.Errorf("An error occured sending request: %s \n", requestErr)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	request.Header.Add("X-API-Version", "1")
	request.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return fmt.Errorf("An error occured getting response: %s \n", responseErr)
	}

	defer response.Body.Close()
	fmt.Println(response.StatusCode)

	switch response.StatusCode {
	case http.StatusOK:
		var respData myTypes.PersonResponse
		if err := utils.ResponseDecoderHandler(response.Body, &respData); err != nil {
			return fmt.Errorf("An error occured while marshalling response body into data: %w", err)
		}
		fmt.Println("✅✅ Successfully fetched data ...")
		fmt.Println(respData)

	case http.StatusUnauthorized:
		// Fetch refresh token
		refreshData, err := utils.RefreshToken(token.RefreshToken)
		if err != nil {
			return fmt.Errorf("failed to get refresh token: %w", err)
		}
		fmt.Println(refreshData)
		return GetProfile(id)
	default:
		var respData myTypes.ErrorResponse
		if err := utils.ResponseDecoderHandler(response.Body, &respData); err != nil {
			return fmt.Errorf("An error occured while marshalling response body into data")
		}
		fmt.Println(respData)
		return fmt.Errorf(" ❎❎ Server returned error of status: %d, and message: %s \n", respData.Status, respData.Message)
	}
	return nil
}
