package pirates

import (
	"fmt"
	"math/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

type Ship struct {
	Captain Pirate
	Crew    []Pirate
}

func (s *Ship) FillShip() {
	s.Captain = Pirate{
		Name:   randSeq(8),
		Parrot: Parrot{COCKATOO, randSeq(5)},
	}
	for i := 0; i < seededRand.Intn(25) + 1; i++ {
		s.Crew = append(s.Crew, Pirate{
			Name:   randSeq(8),
			Parrot: Parrot{CAIQUE, randSeq(5)},
		})
	}
	fmt.Printf("Filled ship, captain is %s, crew number %d\n", s.Captain.Name, len(s.Crew))
}

func (s *Ship) Battle(otherShip *Ship)  {
	firstShipScore := -s.Captain.Intoxication
	for _, pirate := range s.Crew {
		if !pirate.Dead {
			firstShipScore++
		}
	}
	secondShipScore := -otherShip.Captain.Intoxication
	for _, pirate := range otherShip.Crew {
		if !pirate.Dead {
			secondShipScore++
		}
	}

	var casualty int
	if firstShipScore > secondShipScore {
		size := len(otherShip.Crew)
		casualty = seededRand.Intn(size) + 1
		fmt.Printf("First ship won, second ship crew size %d, casualty %d\n", size, casualty)
		otherShip.Crew = otherShip.Crew[:size-casualty]
		fmt.Printf("Second ship crew size %d\n", len(otherShip.Crew))
	} else {
		size := len(s.Crew)
		casualty = seededRand.Intn(size) + 1
		fmt.Printf("Second ship won, first ship crew size %d, casualty %d\n", size, casualty)
		s.Crew = s.Crew[:size-casualty]
		fmt.Printf("First ship crew size %d\n", len(s.Crew))
	}

}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
