// Package components serves as a store for all custom components used in the application.
package components

import (
	"fmt"
	"strings"

	"github.com/andjrue/go-postman/internal/collections"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TreeView struct {
	items              []treeItem
	cursor             int
	directoryTextInput textinput.Model
	showDirectoryInput bool
}

type treeItem struct {
	name   string
	isDir  bool
	method string
	indent int
}

type CreateDirectoryMsg struct {
	name string
}

func createDirectoryTextInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Enter a new directory name..."
	ti.CharLimit = 50
	ti.Width = 30

	return ti
}

func (t *TreeView) resetTextInput() {
	t.showDirectoryInput = false
	t.directoryTextInput.SetValue("")
	t.directoryTextInput.Blur()
}

func NewTreeView(coll *collections.Collection) *TreeView {
	dti := createDirectoryTextInput()
	items := []treeItem{}

	for dirName, directory := range *coll {
		items = append(items, treeItem{
			name:   dirName,
			isDir:  true,
			indent: 0,
		})

		for reqName, request := range directory {
			items = append(items, treeItem{
				name:   reqName,
				isDir:  false,
				method: request.Method,
				indent: 1,
			})
		}
	}

	return &TreeView{
		items:              items,
		cursor:             0,
		directoryTextInput: dti,
		showDirectoryInput: false,
	}
}

func (t *TreeView) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	if t.showDirectoryInput {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "enter":
				dirName := t.directoryTextInput.Value()
				t.resetTextInput()
				return func() tea.Msg {
					return CreateDirectoryMsg{name: dirName}
				}

			case "esc":
				t.resetTextInput()
				return nil
			}
		}

		t.directoryTextInput, cmd = t.directoryTextInput.Update(msg)
		return cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if t.cursor > 0 {
				t.cursor--
			}

		case "down", "j":
			if t.cursor < len(t.items)-1 {
				t.cursor++
			}

		case "N":
			t.showDirectoryInput = true
			t.directoryTextInput.Focus()
			return textinput.Blink
		}
	}

	return nil
}

func (t *TreeView) View() string {
	var s string

	if t.showDirectoryInput {
		s += "Create New Directory:\n\n"
		s += t.directoryTextInput.View() + "\n\n"
		s += "(press enter to create, esc to cancel)"

		return s
	}

	s = "Collections:\n\n"

	for i, item := range t.items {
		cursor := " "
		if t.cursor > i {
			cursor = ">"
		}

		indent := strings.Repeat("   ", item.indent)

		if item.isDir {
			s += fmt.Sprintf("%s %s %s\n", cursor, indent, item.name)
		} else {
			s += fmt.Sprintf("%s %s└─ %s [%s]\n", cursor, indent, item.name, item.method)
		}
	}

	s += "\nPress N to create directory, q to quit, up and down arrows to navigate\n"
	return s
}
