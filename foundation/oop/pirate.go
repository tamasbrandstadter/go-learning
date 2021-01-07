package oop

import (
	"fmt"
	"math/rand"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	LOVEBIRD Breed = iota
	CAIQUE
	COCKATOO
)

type Breed byte

type Parrot struct {
	Breed
	Name string
}

type Pirate struct {
	Intoxication int
	PassedOut    bool
	Dead         bool
	Name         string
	Parrot
}

func (p *Pirate) DrinkSomeRum() {
	if !p.PassedOut && !p.Dead {
		p.Intoxication++
		fmt.Printf("%s intox %d\n", p.Name, p.Intoxication)
	} else {
		fmt.Printf("%s is passed out/dead, can't drink more\n", p.Name)
	}
}

func (p *Pirate) HowsItGoingMate() {
	if p.Dead {
		fmt.Printf("%s is dead\n", p.Name)
	} else if p.PassedOut {
		fmt.Printf("%s is passed out\n", p.Name)
	} else {
		switch p.Intoxication {
		case 0, 1, 2, 3, 4:
			fmt.Printf("%s says: Pour me anudder!\n", p.Name)
		default:
			fmt.Println("Arghh, I'ma Pirate. How d'ya d'ink its goin?")
			p.PassedOut = true
		}
	}
}

func (p *Pirate) Die() {
	p.Dead = true
}

func (p *Pirate) Resurrect() {
	p.Dead = false
	p.PassedOut = false
	p.Intoxication = 0
}

func (p *Pirate) WakeUp() {
	if !p.Dead {
		p.PassedOut = false
		p.Intoxication = 0
	}
}

func (p *Pirate) Brawl(otherP *Pirate) {
	if areTheyBrawlReady(p, otherP) {
		i := seededRand.Intn(4-1) + 1
		switch i {
		case 1:
			p.Dead = true
		case 2:
			otherP.Dead = true
		default:
			p.PassedOut = true
			otherP.PassedOut = true
		}
	} else {
		fmt.Println("Pirates can't brawl")
	}
}

func areTheyBrawlReady(p *Pirate, otherP *Pirate) bool {
	return !p.Dead && !otherP.Dead && !p.PassedOut && !otherP.PassedOut
}
