package app

import (
	"fmt"

	"github.com/andjrue/go-postman/internal/collections"
	"github.com/andjrue/go-postman/internal/components"
	tea "github.com/charmbracelet/bubbletea"
)

func Run() {
	coll, err := collections.LoadFile()
	if err != nil {
		fmt.Printf("error loading config: %v", err)
		return
	}

	p := tea.NewProgram(components.NewModel(coll))
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
