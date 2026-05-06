package profile

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/towbee05/insighta/pkg/myTypes"
	"github.com/towbee05/insighta/pkg/utils"
)

func SearchProfile(q string) error {
	token, tokenErr := utils.CheckTokenFileExistence()
	if tokenErr != nil {
		return fmt.Errorf("an error occured generating token: %s \n", tokenErr)
	}
	var url string = fmt.Sprintf("%s/search?q=%s", BASE_URL, url.QueryEscape(q))

	response, respErr := utils.MakeGetRequest(url, *token)
	if respErr != nil {
		return fmt.Errorf("%s", respErr)
	}

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
		return SearchProfile(q)
	default:
		var respData myTypes.ErrorResponse
		if err := utils.ResponseDecoderHandler(response.Body, &respData); err != nil {
			return fmt.Errorf("An error occured while marshalling response body into data: %s", err)
		}
		fmt.Println(respData)
		return fmt.Errorf(" ❎❎ Server returned error of status: %d, and message: %s \n", respData.Status, respData.Message)
	}
	return nil
}
