package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func getHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode > 200 {
		return "", errors.New(fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body))
	}

	if resp.Header.Get("Content-Type") != "text/html" {
		return "", errors.New("Response is not HTML")
	}

	return string(body), nil
}
