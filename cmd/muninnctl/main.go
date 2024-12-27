package main

import (
    "fmt"
    "os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tenebresus/muninn/pkg/muninnctl"
)

func main() {

    muninnModel := muninnctl.InitializeModel()    
    p := tea.NewProgram(muninnModel)

    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }

}
