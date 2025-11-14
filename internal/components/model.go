package components

import (
	"github.com/andjrue/go-postman/internal/collections"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	collection *collections.Collection
	tree *TreeView
}

func NewModel(coll *collections.Collection) Model {
	return Model{
		collection: coll,
		tree: NewTreeView(coll),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	
	m.tree.Update(msg)
	return m, nil
}

func (m Model) View() string {
	return m.tree.View()
}