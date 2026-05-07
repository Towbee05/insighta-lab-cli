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
	var GithubURL string = "https://hng-stage-1-eight-tan.vercel.app/auth/github"
	response, err := http.Get(GithubURL)
	if err != nil {
		return fmt.Errorf("error fetching authentication endpoint: %w", err)
	}
	// close network  connection to prevent data leakage.
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("An error occured. Unexpected status %d", response.StatusCode)
	}

	var data myTypes.GithubAuthData
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return fmt.Errorf("error decoding authentication data: %w", err)
	}

	authentication_url := data.Authentication_url
	state := data.State
	fmt.Println(authentication_url)

	browser.OpenURL(authentication_url)

	return utils.RunningUntilPoll(state)
}
