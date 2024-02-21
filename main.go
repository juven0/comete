package main

import (
	"comete/cmd/tea"
	"comete/internal/types"
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tearoot "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

	columns := []table.Column{
		{Title: "Condition", Width: 10},
		{Title: "Temperature ", Width: 10},
	}

	rows := []table.Row{
		{"tsara", "19 C"},
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

	m := tea.Model{Table: t}

	if _, err := tearoot.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
