package main

import (
	"comete/cmd/tea"
	"fmt"
	"os"

	tearoot "github.com/charmbracelet/bubbletea"
)

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

	// data, err := app.Fetcher("ambositra", "hourly")
	// if err != nil {
	// 	fmt.Println("no data forme ")

	// } else {
	// 	jsonData := bytes.NewReader(data)

	// 	if err := json.NewDecoder(jsonData).Decode(&Weathers); err != nil {
	// 		return
	// 	}
	// }

	model := tea.NewModel()

	if _, err := tearoot.NewProgram(model).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
