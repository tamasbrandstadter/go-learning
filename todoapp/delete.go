package todoapp

import (
	"bufio"
	"log"
	"os"
)

type DeleteCommand struct {
	description string
	flag        string
}

func NewDeleteCommand() Command {
	return &DeleteCommand{
		description: "Delete todo",
		flag:        "-d",
	}
}

func (d DeleteCommand) GetDescription() string {
	return d.description
}

func (d DeleteCommand) GetFlag() string {
	return d.flag
}

func (d DeleteCommand) Execute(num int) {
	open, _ := os.Open(FileName)
	scanner := bufio.NewScanner(open)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for index, line := range lines {
		if index != num-1 {
			if _, err := file.WriteString(line + "\n"); err != nil {
				log.Println(err)
			}
		}
	}
}
