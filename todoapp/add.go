package todoapp

import (
	"log"
	"os"
)

type AddCommand struct {
	description string
	flag        string
}

func NewAddCommand() Command {
	return &AddCommand{
		description: "Add todo",
		flag:        "-a",
	}
}

func (a AddCommand) GetDescription() string {
	return a.description
}

func (a AddCommand) GetFlag() string {
	return a.flag
}

func (a AddCommand) Execute(task string) {
	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	if _, err := file.WriteString("\n" + task); err != nil {
		log.Println(err)
	}
}
