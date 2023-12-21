package cmd

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [link]",
	Short: "Krunch a link",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		link := args[0]

		// Initialize the Bubble Tea program
		p := tea.NewProgram(initialModel())

		// Create channels for signaling completion
		doneTea := make(chan bool)
		doneKrunch := make(chan bool)

		// Run the Bubble Tea program in a separate goroutine
		go func() {
			if _, err := p.Run(); err != nil {
				fmt.Println("Error running Bubble Tea program:", err)
			}
			doneTea <- true
		}()

		// Process the link in a separate goroutine
		go krunchLink(link, doneKrunch)

		// Wait for both the Bubble Tea program and link processing to finish
		<-doneTea
		<-doneKrunch
	},
}

type errMsg error

type model struct {
	spinner  spinner.Model
	quitting bool
	err      error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s}
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Krunching your link\n\n", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}

	return str
}

func krunchLink(link string, done chan<- bool) {
	fmt.Printf("Krunching link: %v\n", link)
	// Simulate link processing
	// ... actual processing logic ...
	done <- true
}

func init() {
	rootCmd.AddCommand(createCmd)
}
