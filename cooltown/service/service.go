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

	/* Putting Audio as json and byte to put into the URI*/
	post, _ := json.Marshal(map[string]string{
		"Audio": mp3,
	})

	query := bytes.NewBuffer(post)

	if req, err := http.NewRequest("POST", SEARCH, query); err == nil {
		if rsp, err := client.Do(req); err == nil {
			defer rsp.Body.Close()
			if rsp.StatusCode == http.StatusOK {
				var response map[string]string
				err = json.NewDecoder(rsp.Body).Decode(&response)
				return response["id"], nil /* Gets the id this is the name of the track */
			}
		}
	}
	return "", errors.New("Service")
}

func ServiceTracks(id string) (string, error) {
	client := &http.Client{}

	/* The id recieved has spaces in the string so replacing the spaces with + as that's how the full audio ids are stored*/
	newId := strings.ReplaceAll(id, " ", "+")

	/* Combine the URI with Id */
	url := TRACKS + newId

	if req, err := http.NewRequest("GET", url, nil); err == nil {
		if rsp, err := client.Do(req); err == nil {
			defer rsp.Body.Close()
			if rsp.StatusCode == http.StatusOK {
				var response map[string]string
				err = json.NewDecoder(rsp.Body).Decode(&response)
				return response["Audio"], nil /* Outputs full audio*/
			}
		}
	}
	return "", errors.New("Service")
}
