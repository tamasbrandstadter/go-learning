package todoapp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const FileName = "todos.csv"

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
		} else if text == "-a" {
			fmt.Print("\nType todo text:\n")
			scanner.Scan()
			task := scanner.Text()
			app.Commands["-a"].(*AddCommand).Execute(task)
		} else if text == "-d" {
			fmt.Print("\nWhich todo do you want to remove? Specify the number.\n")
			scanner.Scan()
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Println("Can't parse the number")
			} else {
				app.Commands["-d"].(*DeleteCommand).Execute(int(num))
			}
		} else {
			fmt.Println("Flag is not supported")
		}
	}

}
