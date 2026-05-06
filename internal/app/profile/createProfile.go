package profile

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/towbee05/insighta/pkg/myTypes"
	"github.com/towbee05/insighta/pkg/utils"
)

var BASE_URL string = "http://localhost:8000/api/profiles/"

func CreateProfile(name string) error {
	token, tokenErr := utils.CheckTokenFileExistence()
	if tokenErr != nil {
		return fmt.Errorf("an error occured generating token: %s \n", tokenErr)
	}
	var jsonBody = myTypes.CreatePerson{
		Name: name,
	}
	marshalledData, marshalledErr := json.Marshal(jsonBody)
	if marshalledErr != nil {
		return fmt.Errorf("failed to marshal data: %s", marshalledErr)
	}

	response, respErr := utils.MakePostRequest(BASE_URL, *token, marshalledData)
	if respErr != nil {
		return fmt.Errorf("%s", respErr)
	}
	fmt.Println(response.StatusCode)

	switch response.StatusCode {
	case http.StatusCreated:
		var respData myTypes.PersonResponse
		if err := utils.ResponseDecoderHandler(response.Body, &respData); err != nil {
			return fmt.Errorf("An error occured while marshalling response body into data: %w", err)
		}
		fmt.Println("✅✅ Successfully created data ...")
		fmt.Println(respData)

	case http.StatusUnauthorized:
		// Fetch refresh token
		refreshData, err := utils.RefreshToken(token.RefreshToken)
		if err != nil {
			return fmt.Errorf("failed to get refresh token: %w", err)
		}
		fmt.Println(refreshData)
		return CreateProfile(name)
	default:
		var respData myTypes.ErrorResponse
		if err := utils.ResponseDecoderHandler(response.Body, &respData); err != nil {
			return fmt.Errorf("An error occured while marshalling response body into data")
		}
		fmt.Println(respData)
		return fmt.Errorf(" ✖️✖️ Server returned error of status: %d, and message: %s \n", respData.Status, respData.Message)
	}
	return nil
}
