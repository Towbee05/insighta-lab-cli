package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/towbee05/insighta/pkg/myTypes"
)

func RequestForToken(state string) (*myTypes.PolledResult, error) {
	url := fmt.Sprintf("http://localhost:8000/auth/github/poll?state=%s", state)

	response, err := http.Get(url)

	fmt.Println(response)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	var data myTypes.PolledResult

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return nil, err
	}

	fmt.Println(data)

	return &data, nil
}
