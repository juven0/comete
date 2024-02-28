package tea

import (
	"comete/internal/types"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var BasStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderBottomForeground(lipgloss.Color("240"))

type Model struct {
	List list.Model
}
type Task struct {
	title       string
	description string
}

func (t Task) FilterValue() string {
	return ""
}

func NewModel() *Model {
	// Item := []list.Item{
	// 	types.Weather{Date: "2024-01-24T05:00:00", Weather: "mostly_cloudy", Summary: "Mostly cloudy", Temperature: 20.5, Cloud_cover: types.Cloud_cover{Total: 75}, Icon: 5, Wind: types.Wind{Speed: 3.4, Angle: 151, Dir: "SSE"}, Precipitation: types.Precipitation{Total: 0, Types: ""}},
	// }
	return &Model{}
}

func (m Model) Init() tea.Cmd { return nil }

func (m *Model) InitList(width, heigth int, title string) {
	// _, err := app.App()
	m.List = list.New([]list.Item{}, list.NewDefaultDelegate(), width, heigth)
	m.List.Title = title
	m.List.SetItems([]list.Item{
		Task{title: "tena mety", description: "qwuowiuoirwego "},
		Task{title: "tena mety", description: "qwuowiuoirwego "},
		types.Weather{Date: "2024-01-24T05:00:00", Weather: "mostly_cloudy", Summary: "Mostly cloudy", Temperature: 20.5, Cloud_cover: types.Cloud_cover{Total: 75}, Icon: 5, Wind: types.Wind{Speed: 3.4, Angle: 151, Dir: "SSE"}, Precipitation: types.Precipitation{Total: 0, Types: ""}},
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.InitList(msg.Width, msg.Height, "weatheres")
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.List.View()
}
