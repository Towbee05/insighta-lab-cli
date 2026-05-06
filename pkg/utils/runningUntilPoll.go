package utils

import (
	"fmt"
	"time"
)

func RunningUntilPoll(state string) error {
	ticker := time.NewTicker(8 * time.Second)
	timeout := time.After(5 * time.Minute)

	defer ticker.Stop()

	fmt.Println("⌛⌛ Waiting for authentication ...")

	for {
		select {
		case <-ticker.C:
			token, err := RequestForToken(state)
			fmt.Println(token)
			if err != nil {
				fmt.Println("retrying...")
			}
			switch token.Status {
			case "pending":
				fmt.Println("⌛ Still authenticating ...")
				continue
			case "error":
				return fmt.Errorf("❌❌ An error occured during authentication %s", *token.Message)
			case "success":
				fmt.Println("Authentication successful. ✅✅")
				return SaveToken(*token.Data)
			}
		case <-timeout:
			return fmt.Errorf("Authentication timed out")
		}
	}
}
