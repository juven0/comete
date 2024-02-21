package tea

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var BasStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderBottomForeground(lipgloss.Color("240"))

type Model struct {
	Table table.Model
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return BasStyle.Render(m.Table.View()) + "\n"
}
