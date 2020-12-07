package main

import (
	"fmt"
	"time"
)

const (
	WHITE = iota
	BLUE
	RED
)

type Color byte

type PostIt struct {
	color Color
	text  string
}

type BlogPost struct {
	authorName      string
	title           string
	text            string
	publicationDate time.Time
}

type Owner struct {
	name string
}

type Animal struct {
	hunger int
	thirst int
	name   string
	Owner
}

func (p PostIt) printText() {
	fmt.Println(p.text)
}

func (bp BlogPost) printDetails() {
	fmt.Printf("This post was created on %s, by %s with title %s and text %s\n", bp.publicationDate.Format(time.RFC3339),
		bp.authorName, bp.title, bp.text)
}

func (a Animal) eat() {
	a.hunger -= 1
}

func (a Animal) drink() {
	a.thirst -= 1
}

func (a Animal) play() {
	fmt.Printf("%s plays with its owner %s\n", a.name, a.Owner.name)
	a.hunger += 1
	a.thirst += 1
}

func (a *Animal) setOwner(newOwner Owner) {
	a.Owner = newOwner
}

func main() {
	red := PostIt{RED, "Idea 1"}
	blue := PostIt{BLUE, "Awesome"}
	white := PostIt{WHITE, "Superb"}

	postIts := [3]PostIt{red, blue, white}

	for _, postIt := range postIts {
		postIt.printText()
	}

	bp := BlogPost{
		authorName:      "alma",
		title:           "title",
		text:            "lorem ipsum",
		publicationDate: time.Now(),
	}
	bp.printDetails()

	cat := Animal{
		hunger: 0,
		thirst: 0,
		name:   "Boxi",
		Owner:  Owner{"John"},
	}

	cat.play()
	cat.setOwner(Owner{"Ed"})
	cat.play()
}
