package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var httpClient = &http.Client{}

func Post(reqUrl string, reqBody map[string]string, contentType string, headers map[string]string) (string, error) {
	requestBody := makeBody(reqBody)
	httpRequest, _ := http.NewRequest("POST", reqUrl, requestBody)

	httpRequest.Header.Add("Content-Type", contentType)
	if headers != nil {
		for k, v := range headers {
			httpRequest.Header.Add(k, v)
		}
	}
	res, err := httpClient.Do(httpRequest)
	if err != nil {
		return "", fmt.Errorf("httpClient.Do error, %w", err)
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code not equal 200, code is %d, %w", http.StatusNotFound, err)
	}
	defer res.Body.Close()
	response, _ := io.ReadAll(res.Body)
	return string(response), nil
}

func makeBody(reqBody map[string]string) io.Reader {
	bytesData, _ := json.Marshal(reqBody)
	return bytes.NewReader(bytesData)
}

func Get(reqUrl string, reqParams map[string]string, contentType string, headers map[string]string) (string, error) {
	urlParams := url.Values{}
	Url, _ := url.Parse(reqUrl)
	for key, val := range reqParams {
		urlParams.Set(key, val)
	}
	urlPath := Url.String()
	httpRequest, _ := http.NewRequest("GET", urlPath, nil)
	if headers != nil {
		for k, v := range headers {
			httpRequest.Header.Add(k, v)
		}
	}
	res, err := httpClient.Do(httpRequest)
	if err != nil {
		return "", fmt.Errorf("httpClient.Do error, %w", err)
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code not equal 200, code is %d, %w", http.StatusNotFound, err)
	}
	defer res.Body.Close()
	response, _ := io.ReadAll(res.Body)
	return string(response), nil
}
