package profile

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"

	"github.com/towbee05/insighta/pkg/myTypes"
	"github.com/towbee05/insighta/pkg/utils"
)

func ExportProfile(filters myTypes.ExportFilters) error {
	var queryParams string = utils.BuildExportParam(filters)
	token, tokenErr := utils.CheckTokenFileExistence()
	if tokenErr != nil {
		return fmt.Errorf("an error occured generating token: %s \n", tokenErr)
	}

	var url string = fmt.Sprintf("%s/export?%s", BASE_URL, queryParams)
	response, respErr := utils.MakeGetRequest(url, *token)
	if respErr != nil {
		return fmt.Errorf("%s", respErr)
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusOK:
		content := response.Header.Get("Content-Disposition")
		var filename string
		_, params, err := mime.ParseMediaType(content)
		if err != nil {
			filename = fmt.Sprintf("profile.%s", filters.Format)
		}
		if contentName, ok := params["filename"]; ok {
			filename = contentName
		}
		filename = fmt.Sprintf("profile.%s", filters.Format)

		file, createFileErr := os.Create(filename)
		if createFileErr != nil {
			return fmt.Errorf("failed to create file: %s", createFileErr)
		}
		defer file.Close()

		if _, err := io.Copy(file, response.Body); err != nil {
			return fmt.Errorf("failed to copy response body into created file: %s", err)
		}

	case http.StatusUnauthorized:
		// Fetch refresh token
		refreshData, err := utils.RefreshToken(token.RefreshToken)
		if err != nil {
			return fmt.Errorf("failed to get refresh token: %w", err)
		}
		fmt.Println(refreshData)
		return ExportProfile(filters)
	default:
		var respData myTypes.ErrorResponse
		if err := utils.ResponseDecoderHandler(response.Body, &respData); err != nil {
			return fmt.Errorf("An error occured while marshalling response body into data")
		}
		fmt.Println(respData)
		return fmt.Errorf(" ✖️✖️ Server returned error of status: %d, and message: %s \n", respData.Status, respData.Message)
	}
	fmt.Println("✅✅ Successfully exported data ...")
	return nil
}
