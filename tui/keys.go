package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit key.Binding
}

var DefaultKeyMap = KeyMap{
	Quit: key.NewBinding(key.WithKeys("ctrl+c"), key.WithHelp("ctrl + c", "Quit")),
}

// TODO: add short and full help helper funcs
