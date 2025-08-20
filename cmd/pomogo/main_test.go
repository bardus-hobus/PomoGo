package main

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestRun(t *testing.T) {
	err := run(tea.WithInput(strings.NewReader("q\n")), tea.WithoutRenderer())
	if err != nil {
		t.Fatalf("run returned error: %v", err)
	}
}

func TestMain(t *testing.T) {
	called := false
	runProgram = func(opts ...tea.ProgramOption) error {
		called = true
		return nil
	}
	main()
	if !called {
		t.Fatalf("expected runProgram to be called")
	}
}
