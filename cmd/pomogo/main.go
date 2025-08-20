package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"pomogo/internal/tui"
)

func run(opts ...tea.ProgramOption) error {
	m := tui.NewMenuModel()
	return tea.NewProgram(m, opts...).Start()
}

var runProgram = run

func main() {
	if err := runProgram(); err != nil {
		log.Fatalf("failed to start program: %v", err)
	}
}
