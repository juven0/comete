package app

import (
	"fmt"
	"net/http"
)

func Fetcher(location string) string {
	cometeUrl := "http://google.com"
	_, err := http.Get(cometeUrl)
	if err != nil {
		fmt.Println("fetche data error !!")
		return ""
	}
	return "data is here ... please change it !!!"
}
