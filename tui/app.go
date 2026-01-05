// Package tui provides the TUI frontend
package tui

import (
	ui "anything-tui/tui/components"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	// UI Components
	HelpView ui.HelpView

	// Keymaps
	keys KeyMap
}

func NewApp() *App {
	return &App{keys: DefaultKeyMap}
}

func (m App) Init() tea.Cmd {
	return nil
}

func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	return m, cmd
}

func (m App) View() string {
	return "hello tui"
}
