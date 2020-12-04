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
	text    string
}

type BlogPost struct {
	authorName string
	title string
	text    string
	publicationDate time.Time
}

func (p PostIt) printText() {
	fmt.Println(p.text)
}

func (bp BlogPost) printDetails() {
	fmt.Printf("This post was created on %s, by %s with title %s and text %s", bp.publicationDate.Format(time.RFC3339),
		bp.authorName, bp.title, bp.text)
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
}
