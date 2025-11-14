// Package components serves as a store for all custom components used in the application.
package components

import (
	"fmt"
	"strings"

	"github.com/andjrue/go-postman/internal/collections"
	tea "github.com/charmbracelet/bubbletea"
)

type TreeView struct {
	items []treeItem
	cursor int
}

type treeItem struct {
	name string
	isDir bool
	method string
	indent int
}

func NewTreeView(coll *collections.Collection) *TreeView {
	items := []treeItem{}
	
	for dirName, directory := range *coll {
		items = append(items, treeItem{
			name: dirName,
			isDir: true, 
			indent: 0,
		})
		
		for reqName, request := range directory {
			items = append(items, treeItem{
				name: reqName, 
				isDir: false,
				method: request.Method,
				indent: 1,
			})
		}
	}
	
	return &TreeView{
		items: items,
	}
}

func (t *TreeView) Update(msg tea.Msg) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
				case "up", "k": 
					if t.cursor > 0 {
						t.cursor--
					}
					
				case "down", "j": 
					if t.cursor < len(t.items) - 1 {
						t.cursor++
					}
			}
	}
}

func (t *TreeView) View() string {
	s := "Collections:\n\n"
	
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
	
	s += "\nPress q to quit, up and down arrows to navigate\n"
	return s
}
