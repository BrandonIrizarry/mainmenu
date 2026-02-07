package mainmenu

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct{}

// Package-local state. I prefer to have this here, rather than stuff
// everything inside the model struct, which I reserve for
// 1. attaching the [tea.Model] interface methods, and 2. exporting
// final state to the model client.
var (
	cursorPos    int
	maxCursorPos int
	choices      []string
)

func (m Model) New(cs []string) {
	choices = cs
}

// The following methods implement the [tea.Model] interface.
func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyEnter:
			// FIXME: the actual model management will
			// start here.
			return m, tea.Quit

		case tea.KeyDown:
			cursorPos++

			// Wrap back to the top if necessary.
			if cursorPos >= maxCursorPos {
				cursorPos = 0
			}

		case tea.KeyUp:
			cursorPos--

			// Wrap back to the bottom if necessary.
			if cursorPos < 0 {
				cursorPos = maxCursorPos
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := strings.Builder{}

	for i := range len(choices) {
		if cursorPos == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(choices[i])
		s.WriteString("\n")
	}

	return s.String()
}
