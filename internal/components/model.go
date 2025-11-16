package components

import (
	"github.com/andjrue/go-postman/internal/collections"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	collection *collections.Collection
	tree       *TreeView
}

func NewModel(coll *collections.Collection) Model {
	return Model{
		collection: coll,
		tree:       NewTreeView(coll),
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

	case CreateDirectoryMsg:
		if msg.name != "" {
			updatedColl, err := collections.AddDirectory(m.collection, msg.name)
			if err == nil {
				m.collection = updatedColl
				m.tree = NewTreeView(m.collection)
			} else {
				// TODO: Display an error message here
				return nil, nil
			}
		}
	}

	cmd := m.tree.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.tree.View()
}
