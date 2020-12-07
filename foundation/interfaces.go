package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Reserve interface {
	PrintDetails()
	generateCode()
}

type TableReservation struct {
	Code       string
	Restaurant string
}

type CarReservation struct {
	Code string
	Car  string
}

func (r *TableReservation) PrintDetails() {
	r.generateCode()
	fmt.Printf("There is a table reservation for %s, with code %s to date %s\n", r.Restaurant, r.Code, getBookingDate())
}

func (r *CarReservation) PrintDetails() {
	r.generateCode()
	fmt.Printf("Car reservation for %s, with code %s to date %s\n", r.Car, r.Code, getBookingDate())
}

func (r *TableReservation) generateCode() {
	code, _ := generateRandomString(8)
	r.Code = code
}

func (r *CarReservation) generateCode() {
	code, _ := generateRandomString(5)
	r.Code = code
}

func getBookingDate() string {
	return time.Now().Format(time.RFC3339)
}

func generateRandomString(n int) (string, error) {
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func main() {
	carReservation := CarReservation{
		Car: "Toyota",
	}
	tableReservation := TableReservation{
		Restaurant: "Test",
	}
	res := make([]Reserve, 2)
	res[0], res[1] = &carReservation, &tableReservation
	for _, r := range res {
		r.PrintDetails()
	}
}
