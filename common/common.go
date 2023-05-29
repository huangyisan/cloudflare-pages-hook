package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendTG(text string, botToken, chatID string) error {
	url := fmt.Sprintf("https://api.telegram.org/%s/sendMessage?chat_id=%s&parse_mode=markdown", botToken, chatID)
	body, _ := json.Marshal(map[string]string{
		"text": text,
	})
	r, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil || r.StatusCode != 200 {
		return fmt.Errorf("tg api status not 200")
	}
	defer r.Body.Close()
	log.Println("send success msg")
	return nil
}
