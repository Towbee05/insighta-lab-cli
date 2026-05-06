package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/towbee05/insighta/pkg/myTypes"
)

func SaveToken(data myTypes.Token) error {
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// Get user config dir
	userConfigDir, conf_err := os.UserConfigDir()
	if conf_err != nil {
		return fmt.Errorf("An error occured locating user config path: %s", conf_err)
	}

	appDir := filepath.Join(userConfigDir, "insighta-lab-cli")

	if err := os.MkdirAll(appDir, 0700); err != nil {
		return fmt.Errorf("An error occured creating config file for insighta lab-cli: %s", conf_err)
	}

	storage := filepath.Join(appDir, "token.json")

	if err := os.WriteFile(storage, res, 0600); err != nil {
		return err
	}
	return nil
}
