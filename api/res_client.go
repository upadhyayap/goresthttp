package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}, headers http.Header) (*http.Response, *error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, &err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBody))
	client := http.Client{}
	res, err := client.Do(request)

	if err != nil {
		return nil, &err
	}

	defer res.Body.Close()

	return res, nil
}
