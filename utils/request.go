package utils

import (
	"io"
	"net/http"
	"time"
)

func MakeRequest(method string, url string, body io.Reader) (*http.Response, error) {
	httpClient := &http.Client{Timeout: 60 * time.Second}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := httpClient.Do(req)
	return res, err
}
