package main

import (
	"bytes"
	"comete/cmd/tea"
	"comete/internal/app"
	"comete/internal/types"
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tearoot "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var Weathers struct {
	Weathers []types.Weather `json:"data"`
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

	// jsonData, err := os.Open("./datatest.json")
	// if err != nil {
	// 	return
	// }
	// defer jsonData.Close()
	data, err := app.Fetcher("ambositra", "hourly")
	if err != nil {
		fmt.Println("no data")
		return
	}
	reader := bytes.NewReader(data)

	if err := json.NewDecoder(reader).Decode(&Weathers); err != nil {
		return
	}

	fmt.Println(Weathers.Weathers[0])

	columns := []table.Column{
		{Title: "icon", Width: 20},
		{Title: "Condition", Width: 10},
		{Title: "Temperature ", Width: 10},
	}

	rows := []table.Row{
		{"", "tsara", "19 C"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
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

	model := tea.Model{Table: t}

	if _, err := tearoot.NewProgram(model).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
