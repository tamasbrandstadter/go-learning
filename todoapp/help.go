package todoapp

import "fmt"

type HelpCommand struct {
	description string
	flag        string
}

func NewHelpCommand() Command {
	return &HelpCommand{
		description: "Print available commands",
		flag:        "-h",
	}
}

func (h HelpCommand) GetDescription() string {
	return h.description
}

func (h HelpCommand) GetFlag() string {
	return h.flag
}

func (h HelpCommand) Execute(commands map[string]Command) {
	fmt.Println("Command Line Todo application\n=============================\nAvailable commands:")
	for _, command := range commands {
		fmt.Println(command.GetFlag() + " " + command.GetDescription())
	}
}
