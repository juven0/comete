package tea

// import (
// 	"comete/internal/app"
// 	"comete/internal/types"
// 	"sync"

// 	"github.com/charmbracelet/bubbles/list"
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// )

// var BasStyle = lipgloss.NewStyle().
// 	BorderStyle(lipgloss.NormalBorder()).
// 	BorderBottomForeground(lipgloss.Color("240"))

// type Model struct {
// 	List list.Model
// }
// type Task struct {
// 	title       string
// 	description string
// }

// var wg sync.WaitGroup

// func (t Task) FilterValue() string {
// 	return ""
// }

// func NewModel() *Model {
// 	Item := []list.Item{
// 		types.Weather{Date: "2024-01-24T05:00:00", Weather: "mostly_cloudy", Summary: "Mostly cloudy", Temperature: 20.5, Cloud_cover: types.Cloud_cover{Total: 75}, Icon: 5, Wind: types.Wind{Speed: 3.4, Angle: 151, Dir: "SSE"}, Precipitation: types.Precipitation{Total: 0, Types: ""}},
// 	}
// 	return &Model{
// 		List: list.NewModel(Item, list.NewDefaultDelegate(), 100, 100),
// 	}
// }

// // func FecthChan(ch chan []types.Weather) {
// // 	weatheres, err := app.App()
// // 	if err != nil {
// // 		return
// // 	}
// // 	ch <- weatheres

// // }

// func (m Model) Init() tea.Cmd { return nil }

// func (m *Model) InitList(width, heigth int, title string) {
// 	ch := make(chan []types.Weather)
// 	go app.App(ch)
// 	//weatheres := <-ch
// 	// m.List = list.New([]list.Item{}, list.NewDefaultDelegate(), width, heigth)
// 	m.List.Title = title
// 	// if len(weatheres) == 0 {
// 	// 	for i, e := range weatheres {
// 	// 		m.List.SetItem(i, e)
// 	// 	}
// 	// }

// }

// func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd
// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		m.InitList(msg.Width, msg.Height, "weatheres")
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "q", "ctrl+c":
// 			return m, tea.Quit
// 		}
// 	}
// 	m.List, cmd = m.List.Update(msg)
// 	return m, cmd
// }

// func (m Model) View() string {
// 	return m.List.View()
// }

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Model struct {
	Table table.Model
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.Table.Focused() {
				m.Table.Blur()
			} else {
				m.Table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return baseStyle.Render(m.Table.View()) + "\n"
}
