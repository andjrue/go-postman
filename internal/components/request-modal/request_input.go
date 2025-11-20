package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
)

type CreateRequestInput struct {
	RequestName string
}

func CreateRequestTextInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "Enter a new request name..."
	ti.CharLimit = 75
	ti.Width = 50

	return ti
}

// this will eventually be a real function
func PrintRequestName(name string) {
	fmt.Printf("request name received: %v", name)
}
