package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

const (
	wordLength = 5
	maxGuesses = 6
)

func main() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

type errMsg error

// model contains the program state and implements the tea.Model interface.
type model struct {
	textInput textinput.Model
	guesses   []string
	id        int
	answer    string
	err       error
}

func initialModel() model {
	// Create text input and set default values.
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = wordLength
	ti.Width = wordLength
	id := getID()

	// Try to use the first command line argument instead.
	return model{
		textInput: ti,
		id:        id,
		answer:    answers[id],
	}
}

// getID returns a wordle ID. If first command line argument is a number use
// that; otherwise get it from the current local date.
func getID() int {
	if len(os.Args) > 1 {
		id, err := strconv.Atoi(os.Args[1])
		if err == nil {
			return id
		}
	}

	now := time.Now()
	tz := now.Location()
	start := time.Date(2021, time.June, 19, 0, 0, 0, 0, tz)
	id := int(now.Sub(start).Hours() / 24)

	// Handle dates before the start and after the end.
	if id < 0 {
		id = -id
	}
	for id >= len(answers) {
		id -= len(answers)
	}
	return id
}

// Init is called before the main game loop.
func (m model) Init() tea.Cmd {
	return textinput.Blink
}

// Update is called when a message is revieved in order to update the model and
// or send command(s).
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m = guess(m)

			// If all guesses were used up or the last guess was the correct
			// answer exit the game.
			if len(m.guesses) == maxGuesses || m.guesses[len(m.guesses)-1] == m.answer {
				result(m)
				return m, tea.Quit
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func guess(m model) model {
	g := m.textInput.Value()

	// Check if the guess is a valid length.
	if len(g) != wordLength {
		return m
	}

	// Check if the guess in in the allowedWords list.
	if _, ok := allowedWords[g]; !ok {
		return m
	}

	m.guesses = append(m.guesses, g)
	m.textInput.SetValue("")
	return m
}

// View renders the program's UI, which is just a string. The view is rendered
// after every Update.
func (m model) View() string {
	var view strings.Builder

	// Render each previous guess.
	unused := "abcdefghijklmnopqrstuvwxyz"
	for i, guess := range m.guesses {
		// Number prefix.
		view.WriteString(strconv.Itoa(i + 1))
		view.WriteString(" ")

		var word strings.Builder
		for ii, c := range guess {
			// Color the letters and list unused letters. Green for correct
			// letter in correct location; yellow for correct letter, but wrong
			// location. Grey for the unused letters list.
			letter := termenv.String(string(c))

			if strings.ContainsRune(m.answer, c) {
				if m.answer[ii] == byte(c) {
					letter = letter.Foreground(termenv.ANSIBlack)
					letter = letter.Background(termenv.ANSIBrightGreen)
				} else {
					letter = letter.Foreground(termenv.ANSIBlack)
					letter = letter.Background(termenv.ANSIBrightYellow)
				}
			}

			// Remove the current character from the unused letter list.
			unused = strings.ReplaceAll(unused, string(c), "")

			word.WriteString(letter.String())
		}
		view.WriteString(word.String())
		view.WriteString("  ")
		view.WriteString(
			termenv.String(unused).Foreground(
				termenv.ANSIBrightBlack,
			).String(),
		)
		view.WriteString("\n")
	}

	// Render the current input.
	view.WriteString(m.textInput.View())

	return view.String()
}

func result(m model) {
	var b strings.Builder
	b.WriteString("Wordle " + strconv.Itoa(m.id))
	if m.guesses[len(m.guesses)-1] == m.answer {
		b.WriteString(" " + strconv.Itoa(len(m.guesses)))
	} else {
		b.WriteString(" X") // Failed to guess the answer.
	}
	b.WriteString("/" + strconv.Itoa(maxGuesses) + "\n\n")

	for _, guess := range m.guesses {
		var word strings.Builder
		for ii, c := range guess {
			if strings.ContainsRune(m.answer, c) {
				if m.answer[ii] == byte(c) {
					word.WriteString("🟩")
				} else {
					word.WriteString("🟨")
				}
			} else {
				word.WriteString("⬛")
			}
		}
		b.WriteString(word.String() + "\n")
	}
	clipboard.WriteAll(b.String())
}
