package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/browser"
	"github.com/towbee05/insighta/pkg/myTypes"
	"github.com/towbee05/insighta/pkg/utils"
)

func Login() error {
	response, err := http.Get("http://localhost:8000/auth/github")
	if err != nil {
		return fmt.Errorf("error fetching authentication endpoint: %w", err)
	}
	// close network  connection to prevent data leakage.
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("An error occured. Unexpected status %d", response.StatusCode)
	}

	// Create a  dict "data" of key "string" and value "string"
	var data myTypes.GithubAuthData
	// Decode response body (it is in byte of data) into  readable hash map format
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return fmt.Errorf("error decoding authentication data: %w", err)
	}

	authentication_url := data.Authentication_url
	state := data.State

	browser.OpenURL(authentication_url)

	return utils.RunningUntilPoll(state)
}
