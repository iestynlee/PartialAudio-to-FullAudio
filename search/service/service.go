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

	/* JSON string converted into a byte array */
	post, _ := json.Marshal(map[string]string{
		"api_token": KEY,
		"audio":     mp3,
	})

	/* This will be added to the end of the URI*/
	query := bytes.NewBuffer(post)

	if req, err := http.NewRequest("POST", URI, query); err == nil {
		if rsp, err := client.Do(req); err == nil {
			defer rsp.Body.Close()
			if rsp.StatusCode == http.StatusOK {
				/* Decoding the response body into json*/
				a := map[string]interface{}{}
				if err := json.NewDecoder(rsp.Body).Decode(&a); err == nil {
					if titles, err := titleOf(a); err == nil {
						return titles, nil /* Success */
					}
				}
			}
		}
	}
	return "", errors.New("Service") /* Whenever there is an error it will be called Service */
}

func titleOf(a map[string]interface{}) (string, error) {
	/* If ok then it will get the result then further into the json there is title and gets the value of title*/
	if titles, ok := a["result"].(map[string]interface{})["title"].(string); ok {
		return titles, nil
	}
	return "", errors.New("titleOf")
}
