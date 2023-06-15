package tools

import (
	"net/http"
)

func MakeRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+AccessToken)
	req.Header.Set("Access-Control-Allow-Origin", "*")

	client := &http.Client{}
	return client.Do(req)
}
