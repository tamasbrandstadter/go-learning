package todoapp

import "fmt"

type HelpCommand struct {
	applicationCommands []Command
	description         string
	flag                string
}

func NewHelpCommand(commands []Command) Command {
	return &HelpCommand{
		applicationCommands: commands,
		description:         "Print help",
		flag:                "-h",
	}
}

func (h HelpCommand) GetDescription() string {
	return h.description
}

func (h HelpCommand) GetFlag() string {
	return h.flag
}

func (h HelpCommand) Execute() {
	fmt.Println("Command Line Todo application\n=============================\nAvailable commands:")
	for _, command := range h.applicationCommands {
		fmt.Println(command.GetFlag() + " " + command.GetDescription())
	}
	fmt.Println(h.flag + " " + h.description)
}
