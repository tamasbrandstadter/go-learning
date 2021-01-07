package todoapp

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type ListCommand struct {
	description string
	flag        string
}

func NewListCommand() Command {
	return &ListCommand{
		description: "List todos",
		flag:        "-l",
	}
}

func (l ListCommand) GetDescription() string {
	return l.description
}

func (l ListCommand) GetFlag() string {
	return l.flag
}

func (l ListCommand) Execute() {
	file, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line++
		fmt.Printf("%d - %s\n", line, scanner.Text())
	}
	if line == 0 {
		fmt.Println("No todos for today!")
	}
}
