package main

import (
	"comete/internal/app"
	"comete/internal/types"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aquasecurity/table"
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

	// model := tea.NewModel()

	// if _, err := tearoot.NewProgram(model).Run(); err != nil {
	// 	fmt.Println("Error running program:", err)
	// 	os.Exit(1)
	// }

	ch := make(chan []types.Weather)
	go app.App(ch)
	weatheres := <-ch

	tb := table.New(os.Stdout)
	tb.SetRowLines(false)
	// tb.SetHeaders("comete")
	// tb.SetHeaderColSpans(0, 6)
	tb.SetHeaderStyle(table.StyleBold)
	tb.SetHeaderStyle(table.StyleBlue)
	tb.SetDividers(table.UnicodeRoundedDividers)
	for _, e := range weatheres {
		//tb.AddHeaders()
		tb.AddRows()
		//tb.AddRow(fmt.Sprintf("%v", e.Date), fmt.Sprintf("%d", e.Icon), e.Weather, e.Summary, fmt.Sprintf("%f", e.Temperature), fmt.Sprintf("%d", e.Cloud_cover.Total))
		jsonData, err := json.Marshal(e)
		if err != nil {
			return
		}
		var data map[string]interface{}
		err = json.Unmarshal(jsonData, &data)
		if err != nil {
			return
		}
		// for _, value := range data {
		// 	fmt.Println(value)
		// }
	}

	tb.Render()
	// tb.Render()

}

func String(i int) {
	panic("unimplemented")
}
