package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/towbee05/insighta/pkg/myTypes"
)

func CheckTokenFileExistence() (*myTypes.Token, error) {
	configDir, confErr := os.UserConfigDir()

	if confErr != nil {
		return nil, confErr
	}

	tokenDir := filepath.Join(configDir, "insighta-lab-cli/token.json")

	var data myTypes.Token

	if _, fileErr := os.Stat(tokenDir); errors.Is(fileErr, os.ErrNotExist) {
		fmt.Println("User session timed out.")
		fmt.Println("Run insighta login to authenticate again. ")
		return nil, fileErr
	} else if fileErr != nil {
		fmt.Println("Unable to access token directory")
		return nil, fileErr
	}

	content, readErr := os.ReadFile(tokenDir)

	if readErr != nil {
		return nil, readErr
	}

	if err := json.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
