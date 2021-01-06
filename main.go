package main

import "./pirates"
import "./todoapp"

func main() {
	pirate := pirates.Pirate{
		Name:   "Silver-Tongue",
		Parrot: pirates.Parrot{Breed: pirates.LOVEBIRD, Name: "Picasso"},
	}
	pirate2 := pirates.Pirate{
		Name:   "Seadog",
		Parrot: pirates.Parrot{Breed: pirates.CAIQUE, Name: "Cayenne"},
	}

	pirate.HowsItGoingMate()
	pirate.DrinkSomeRum()

	pirate2.HowsItGoingMate()
	pirate2.DrinkSomeRum()

	pirate.Brawl(&pirate2)

	pirate.HowsItGoingMate()
	pirate.DrinkSomeRum()
	pirate2.HowsItGoingMate()
	pirate2.DrinkSomeRum()

	ship := pirates.Ship{}
	ship.FillShip()

	ship2 := pirates.Ship{}
	ship2.FillShip()

	ship.Battle(&ship2)

	help := todoapp.NewHelpCommand()
	list := todoapp.NewListCommand()

	commands := make(map[string]todoapp.Command)
	commands[help.GetFlag()] = help
	commands[list.GetFlag()] = list

	app := todoapp.Application{Commands: commands}
	app.Run()
}
