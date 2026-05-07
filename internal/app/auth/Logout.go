package auth

import (
	"fmt"
	"os"
	"path/filepath"
)

func Logout() error {
	// get user config dir
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("An error occured locating user config path: %s", err)
	}

	tokenDir := filepath.Join(userConfigDir, "insighta-lab-cli/token.json")

	if err := os.Remove(tokenDir); err != nil {
		return fmt.Errorf("An error occured writting into token dir: %s", err)
	}
	fmt.Println("Successfully Logged out...")
	return nil
}
