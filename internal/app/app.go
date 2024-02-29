package app

import (
	"comete/internal/types"
	"encoding/json"
	"fmt"
	"os"
)

var Weathers struct {
	Weathers []types.Weather `json:"data"`
}

func App(ch chan []types.Weather) {
	jsonData, err := os.Open("./datatest.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := json.NewDecoder(jsonData).Decode(&Weathers); err != nil {
		fmt.Println(err)
		return
	}
	ch <- Weathers.Weathers
	return
}
func GetLocation() string {
	fmt.Println("do some things....")
	//please complet this ....!!
	return ""
}
