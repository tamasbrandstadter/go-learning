package main

import "./foundation/oop/pirates"
import "./foundation/oop/todoapp"

func main() {
	ship := pirates.Ship{}
	ship.FillShip()
	ship2 := pirates.Ship{}
	ship2.FillShip()

	ship.Battle(&ship2)

	list := todoapp.NewListCommand()
	quit := todoapp.NewQuitCommand()
	add := todoapp.NewAddCommand()
	deleteCommand := todoapp.NewDeleteCommand()
	help := todoapp.NewHelpCommand([]todoapp.Command{list, quit, add, deleteCommand})

	commands := make(map[string]todoapp.Command)
	commands[help.GetFlag()] = help
	commands[list.GetFlag()] = list
	commands[quit.GetFlag()] = quit
	commands[add.GetFlag()] = add
	commands[deleteCommand.GetFlag()] = deleteCommand

	app := todoapp.Application{Commands: commands}
	app.Run()
}
