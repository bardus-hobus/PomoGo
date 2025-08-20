package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewMenuModel(t *testing.T) {
	m := NewMenuModel()
	if len(m.choices) != 2 {
		t.Fatalf("expected 2 choices, got %d", len(m.choices))
	}
}

func TestMenuNavigation(t *testing.T) {
	m := NewMenuModel()
	// Move down to 'Exit'
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = updated.(MenuModel)
	if m.cursor != 1 {
		t.Fatalf("expected cursor 1, got %d", m.cursor)
	}
	// Move up back to 'Start Timer'
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyUp})
	m = updated.(MenuModel)
	if m.cursor != 0 {
		t.Fatalf("expected cursor 0, got %d", m.cursor)
	}
}

func TestMenuExit(t *testing.T) {
	m := NewMenuModel()
	// Move to 'Exit' and select it
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = updated.(MenuModel)
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("expected a command")
	}
	if cmd() != tea.Quit() {
		t.Fatalf("expected tea.Quit command")
	}
	m = updated.(MenuModel)
	if !m.quitting {
		t.Fatal("expected model to be quitting")
	}
}

func TestMenuStartTimer(t *testing.T) {
	m := NewMenuModel()
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = updated.(MenuModel)
	if cmd != nil {
		t.Fatalf("expected nil command, got %v", cmd)
	}
	if m.quitting {
		t.Fatal("start timer should not quit")
	}
}

func TestMenuView(t *testing.T) {
	m := NewMenuModel()
	view := m.View()
	if !strings.Contains(view, "Start Timer") || !strings.Contains(view, "Exit") {
		t.Fatalf("view missing expected choices: %s", view)
	}
	// Ensure goodbye message appears when quitting
	m.quitting = true
	view = m.View()
	if !strings.Contains(view, "Goodbye!") {
		t.Fatalf("expected goodbye message, got %s", view)
	}
}

func TestMenuQuitKeys(t *testing.T) {
	m := NewMenuModel()
	// Test quitting with 'q'
	updated, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	m = updated.(MenuModel)
	if cmd == nil || cmd() != tea.Quit() || !m.quitting {
		t.Fatalf("expected quit command with 'q'")
	}
	// Test quitting with ctrl+c
	m = NewMenuModel()
	updated, cmd = m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	m = updated.(MenuModel)
	if cmd == nil || cmd() != tea.Quit() || !m.quitting {
		t.Fatalf("expected quit command with ctrl+c")
	}
}
