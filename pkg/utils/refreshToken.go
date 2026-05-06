package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/towbee05/insighta/pkg/myTypes"
)

func RefreshToken(refreshToken string) (*myTypes.Token, error) {
	// send RefreshToken
	body := myTypes.RefreshTokenBody{
		RefreshToken: refreshToken,
	}
	marshalledData, marshalErr := json.Marshal(body)
	if marshalErr != nil {
		return nil, fmt.Errorf("failed to convert body to byte %w", marshalErr)
	}
	request, requestErr := http.NewRequest("POST", "http://localhost:8000/auth/refresh", bytes.NewBuffer(marshalledData))
	if requestErr != nil {
		return nil, fmt.Errorf("Could not request for new token: %w", requestErr)
	}
	client := http.Client{}
	response, respErr := client.Do(request)
	if respErr != nil {
		return nil, fmt.Errorf("failed to fetch response: %w", respErr)
	}

	defer response.Body.Close()
	var data myTypes.Token
	switch response.StatusCode {
	case http.StatusCreated:
		if err := ResponseDecoderHandler(response.Body, &data); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		// save refresh token
		if err := SaveToken(data); err != nil {
			return nil, fmt.Errorf("failed to save token: %w", err)
		}
		return &data, nil
	case http.StatusUnauthorized:
		var unAuthorizedResponse myTypes.ErrorResponse
		if err := ResponseDecoderHandler(response.Body, &unAuthorizedResponse); err != nil {
			return nil, fmt.Errorf("An error occured while marshalling response body into data: %w", err)
		}
		fmt.Println(unAuthorizedResponse)
		return nil, fmt.Errorf(" ❎❎ You need to login again: %s\n", unAuthorizedResponse.Message)
	default:
		var errorData myTypes.ErrorResponse
		if err := ResponseDecoderHandler(response.Body, &errorData); err != nil {
			return nil, fmt.Errorf("An error occured while marshalling response body into data: %w", err)
		}
		fmt.Println(errorData)
		return nil, fmt.Errorf("failed to create token: %s ", errorData.Message)
	}
}
