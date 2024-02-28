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

func App() ([]types.Weather, error) {
	jsonData, err := os.Open("./datatest.json")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err := json.NewDecoder(jsonData).Decode(&Weathers); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return Weathers.Weathers, nil
}
func GetLocation() string {
	fmt.Println("do some things....")
	//please complet this ....!!
	return ""
}
