package profile

import (
	"fmt"
	"net/http"

	"github.com/towbee05/insighta/pkg/myTypes"
	"github.com/towbee05/insighta/pkg/utils"
)

func ListResponse(filters myTypes.Filters) error {
	token, tokenErr := utils.CheckTokenFileExistence()
	if tokenErr != nil {
		return fmt.Errorf("an error occured generating token: %s \n", tokenErr)
	}

	params := utils.BuildQueryParams(filters)
	url := fmt.Sprintf("%s?%s", BASE_URL, params)
	response, respErr := utils.MakeGetRequest(url, *token)
	if respErr != nil {
		return fmt.Errorf("%s", respErr)
	}
	fmt.Println(response.StatusCode)

	switch response.StatusCode {
	case http.StatusOK:
		var respData myTypes.ProfileResponse
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
		return ListResponse(filters)
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
