package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
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

var httpClient = &http.Client{}

func Post(url string, reqParams map[string]string, contentType string, headers map[string]string) (string, error) {
	requestBody := makeBody(reqParams)
	httpRequest, _ := http.NewRequest("POST", url, requestBody)

	httpRequest.Header.Add("Content-Type", contentType)
	if headers != nil {
		for k, v := range headers {
			httpRequest.Header.Add(k, v)
		}
	}
	res, err := httpClient.Do(httpRequest)
	if err != nil {
		return "", errors.Wrap(err, "httpClient.Do error")
	}
	if res.StatusCode != 200 {
		return "", errors.Wrap(fmt.Errorf("status code not equal 200, code is %d", http.StatusNotFound), "")
	}
	defer res.Body.Close()
	response, _ := io.ReadAll(res.Body)
	return string(response), nil
}

func makeBody(reqParams map[string]string) io.Reader {
	bytesData, _ := json.Marshal(reqParams)
	return bytes.NewReader(bytesData)
}
