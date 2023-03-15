package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const (
	URI = "https://api.audd.io/recognize"
	KEY = "0d6fa17de127aa8661283db22679380c"
)

func Service(mp3 string) (string, error) {
	client := &http.Client{}

	postBody, _ := json.Marshal(map[string]string{
		"api_token": KEY,
		"audio":     mp3,
	})

	responseBody := bytes.NewBuffer(postBody)

	if req, err := http.NewRequest("POST", URI, responseBody); err == nil {
		if rsp, err := client.Do(req); err == nil {
			defer rsp.Body.Close()
			if rsp.StatusCode == http.StatusOK {
				a := map[string]interface{}{}
				if err := json.NewDecoder(rsp.Body).Decode(&a); err == nil {
					if titles, err := titleOf(a); err == nil {
						return titles, nil
					}
				}
			}
		}
	}
	return "", errors.New("Service")
}

func titleOf(a map[string]interface{}) (string, error) {
	if titles, ok := a["result"].(map[string]interface{})["title"].(string); ok {
		return titles, nil
	}
	return "", errors.New("titleOf")
}
