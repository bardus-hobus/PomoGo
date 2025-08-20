package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// MenuModel represents the state of the main menu.
type MenuModel struct {
	cursor   int
	choices  []string
	quitting bool
}

// NewMenuModel creates a new menu model with default choices.
func NewMenuModel() MenuModel {
	return MenuModel{
		choices: []string{"Start Timer", "Exit"},
	}
}

// Init initializes the menu model.
func (m MenuModel) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model's state.
func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			switch m.cursor {
			case 0:
				// Start Timer - not yet implemented
			case 1:
				m.quitting = true
				return m, tea.Quit
			}
		}
	}
	return m, nil
}

// View renders the menu.
func (m MenuModel) View() string {
	var b strings.Builder
	b.WriteString("Select an option:\n\n")
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		fmt.Fprintf(&b, "%s %s\n", cursor, choice)
	}
	if m.quitting {
		b.WriteString("\nGoodbye!\n")
	}
	return b.String()
}
