package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	SEARCH = "http://localhost:3001/search"
	TRACKS = "http://localhost:3000/tracks/"
)

func ServiceAudio(mp3 string) (string, error) {
	client := &http.Client{}

	query, _ := json.Marshal(map[string]string{
		"Audio": mp3,
	})

	responseBody := bytes.NewBuffer(query)

	if req, err := http.NewRequest("POST", SEARCH, responseBody); err == nil {
		if rsp, err := client.Do(req); err == nil {
			defer rsp.Body.Close()
			if rsp.StatusCode == http.StatusOK {
				var response map[string]string
				err = json.NewDecoder(rsp.Body).Decode(&response)
				return response["id"], nil
			}
		}
	}
	return "", errors.New("Service")
}

func ServiceTracks(id string) (string, error) {
	client := &http.Client{}

	newId := strings.ReplaceAll(id, " ", "+")

	url := TRACKS + newId

	if req, err := http.NewRequest("GET", url, nil); err == nil {
		if rsp, err := client.Do(req); err == nil {
			defer rsp.Body.Close()
			if rsp.StatusCode == http.StatusOK {
				var response map[string]string
				err = json.NewDecoder(rsp.Body).Decode(&response)
				return response["Audio"], nil
			}
		}
	}
	return "", errors.New("Service")
}
