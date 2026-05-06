package utils

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/towbee05/insighta/pkg/myTypes"
)

func MakePostRequest(url string, token myTypes.Token, body []byte) (*http.Response, error) {
	request, requestErr := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if requestErr != nil {
		return nil, fmt.Errorf("An error occured sending request: %s \n", requestErr)
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	request.Header.Add("X-API-Version", "1")
	request.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	response, responseErr := client.Do(request)
	if responseErr != nil {
		return nil, fmt.Errorf("An error occured getting response: %s \n", responseErr)
	}

	defer response.Body.Close()
	return response, nil
}
