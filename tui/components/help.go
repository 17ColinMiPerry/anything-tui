// Package components provides the TUI ui components
package components

import (
	tea "github.com/charmbracelet/bubbletea"
)

type HelpView struct{}

func NewHelpView() *HelpView {
	return &HelpView{}
}

func (h HelpView) Init() tea.Cmd {
	return nil
}

func (h HelpView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return h, nil
}

func (h HelpView) View() string {
	return ""
}
