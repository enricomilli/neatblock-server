package msg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func MsgTelegram(msg string) error {

	url := "https://api.telegram.org/bot" + os.Getenv("TELEGRAM_BOT_API") + "/sendMessage"
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	payload := map[string]string{
		"chat_id": chatID,
		"text":    msg,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error sending HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
