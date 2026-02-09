package mainmenu

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	cursorPos    int
	maxCursorPos int
	choices      []string
	models       []tea.Model
}

func New() Model {
	choices := []string{
		"Chat UI",
		"Select assets (files, project directories)",
		"Select LLM model",
		"Exit",
	}

	return Model{
		choices:      choices,
		maxCursorPos: len(choices) - 1,
	}
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
			if m.cursorPos == m.maxCursorPos {
				return m, tea.Quit
			}

		case tea.KeyDown:
			m.cursorPos++

			// Wrap back to the top if necessary.
			if m.cursorPos > m.maxCursorPos {
				m.cursorPos = 0
			}

		case tea.KeyUp:
			m.cursorPos--

			// Wrap back to the bottom if necessary.
			if m.cursorPos < 0 {
				m.cursorPos = m.maxCursorPos
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	s := strings.Builder{}

	for i := range len(m.choices) {
		if m.cursorPos == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.choices[i])
		s.WriteString("\n")
	}

	return s.String()
}
