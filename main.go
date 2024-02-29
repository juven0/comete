package main

import (
	"comete/cmd/tea"
	"comete/internal/app"
	"comete/internal/types"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tearoot "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 10},
	}

	rows := []table.Row{}
	for _, e := range weatheres {
		rows = append(rows, e)
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := tea.Model{Table: t}
	if _, err := tearoot.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
