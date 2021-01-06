package todoapp

import (
	"bufio"
	"fmt"
	"os"
)

type Command interface {
	GetDescription() string
	GetFlag() string
}

type Application struct {
	running  bool
	Commands map[string]Command
}

func (app Application) Run() {
	app.running = true

	fmt.Println("Welcome to command line todo app! Use -h for printing help.")

	scanner := bufio.NewScanner(os.Stdin)
	for app.running == true {
		fmt.Print("\nType a command flag:\n")
		scanner.Scan()
		text := scanner.Text()
		if text == "-l" {
			app.Commands[text].(*ListCommand).Execute()
		} else if text == "-h" {
			app.Commands[text].(*HelpCommand).Execute()
		} else if text == "-q" {
			app.Commands[text].(*QuitCommand).Execute(&app)
		} else {
			fmt.Println("Flag is not supported")
		}
	}

}
