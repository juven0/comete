package app

import (
	"io"
	"net/http"
	"os"
)

var BASE_URL = os.Getenv("BASE_URL")

func Fetcher(location, mode string) ([]byte, error) {
	cometeUrl := "https://comete.onrender.com/api/comete/ambositra/day"
	resp, err := http.Get(cometeUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}
