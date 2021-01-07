package todoapp

import "fmt"

type QuitCommand struct {
	description string
	flag        string
}

func NewQuitCommand() Command {
	return &QuitCommand{
		description: "Exit the application",
		flag:        "-q",
	}
}

func (q QuitCommand) GetDescription() string {
	return q.description
}

func (q QuitCommand) GetFlag() string {
	return q.flag
}

func (q QuitCommand) Execute(app *Application) {
	fmt.Println("Quitting application...")
	app.running = false
}
