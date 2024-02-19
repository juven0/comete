package main

import (
	"comete/internal/types"
	"encoding/json"
	"fmt"
	"os"
)

type Weathers struct {
	Weathers []types.Weather `json:"weathers"`
}

func main() {

	asciiArt := `
                                              $$\               
                                              $$ |              
 $$$$$$$\  $$$$$$\  $$$$$$\$$$$\   $$$$$$\  $$$$$$\    $$$$$$\  
$$  _____|$$  __$$\ $$  _$$  _$$\ $$  __$$\ \_$$  _|  $$  __$$\ 
$$ /      $$ /  $$ |$$ / $$ / $$ |$$$$$$$$ |  $$ |    $$$$$$$$ |
$$ |      $$ |  $$ |$$ | $$ | $$ |$$   ____|  $$ |$$\ $$   ____|
\$$$$$$$\ \$$$$$$  |$$ | $$ | $$ |\$$$$$$$\   \$$$$  |\$$$$$$$\ 
 \_______| \______/ \__| \__| \__| \_______|   \____/  \_______|
                                                                
                                                                `
	fmt.Println(asciiArt)
	fmt.Println("☀️")

	jsonData, err := os.Open("./datatest.json")
	if err != nil {
		return
	}
	defer jsonData.Close()
	var weathers Weathers
	data := json.NewDecoder(jsonData)
	data.Decode(&weathers)
	fmt.Println(weathers.Weathers[0])

}
