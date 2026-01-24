package app

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() error {
	m := NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	logger, _ := tea.LogToFile("debug.log", "debug")
	defer logger.Close()

	_, err := p.Run()
	return err
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

type Model struct {
	appName string
}

func NewModel() *Model {
	return &Model{
		appName: "xytz",
	}
}
