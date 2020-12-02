package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

const totalSeconds int = 24 * 60 * 60

type filterFunc func(int) bool

func printNames(names ...string) {
	for index := range names {
		fmt.Println(names[index])
	}
}

func printDetails(name string, age int, height float64, married bool) {
	maritalStatus := "single"
	if married {
		maritalStatus = "married"
	}
	fmt.Printf("%s is %d years old and %.2f tall, %s\n", name, age, height, maritalStatus)
}

func calculateDivision(a, b int) int {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("Can't divide by 0!!")
		}
	}()
	return a / b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func calculateAvgAndSum() (float64, float64) {
	scanner := bufio.NewScanner(os.Stdin)
	var sum float64
	var count int

	fmt.Print("Enter a number or blank to exit:\n")
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			break
		}

		num, err := strconv.ParseFloat(text, 64)
		if err != nil {
			panic(err)
		}
		sum += num
		count++
	}

	avg := sum / float64(count)
	return avg, sum
}

func calculateWorkingHours(dailyHours float64, semesterWeeks float64, workDays float64) float64 {
	return dailyHours * semesterWeeks * workDays
}

func swapVariables(a *int, b *int) (x int, y int) {
	tmp := a
	a = b
	b = tmp
	return *a, *b
}

func remainingSecondsFromDay() int {
	location, err := time.LoadLocation("Local")
	if err != nil {
		panic(err)
	}

	localTime := time.Now().In(location)
	fmt.Println("Local time is " + localTime.String())

	hour, min, sec := localTime.Clock()
	elapsedSeconds := hour*3600 + min*60 + sec

	return totalSeconds - elapsedSeconds
}

func isOdd(num int) bool {
	return num%2 != 0
}

func isEven(num int) bool {
	return num%2 == 0
}

func filterNumbers(numbers []int, f filterFunc) []int {
	var filtered []int
	for i := range numbers {
		num := numbers[i]
		if f(num) {
			filtered = append(filtered, num)
		}
	}
	return filtered
}

func checkDay(day time.Weekday) {
	switch day {
	case time.Monday:
		fmt.Println("It's Monday...")
	case time.Tuesday:
		fmt.Println("Are we there yet?")
	case time.Wednesday, time.Thursday:
		fmt.Println("Almost")
	case time.Friday:
		fmt.Println("Yey")
	default:
		fmt.Println("It's weekend")
	}
}

func multiplicationTable(num int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", i, num, i*num)
	}
}

func countFrom(num, otherNum int) {
	if otherNum <= num {
		fmt.Println("Second number should be bigger")
	}
	for i := num; i < otherNum; i++ {
		fmt.Println(i)
	}
}

func printMatrix() {
	var matrix [4][4]int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == j {
				matrix[i][j] = 1
			}
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	printNames("Esther", "Mary", "Joe")
	printDetails("Tom", 33, 172.5, true)

	calculateDivision(5, 0)
	sum, product := calculateSumAndProduct(2, 6)
	fmt.Printf("Sum=%d, product=%d\n", sum, product)
	avg, sum2 := calculateAvgAndSum()
	fmt.Printf("Avg=%.2f, Sum=%.2f\n", avg, sum2)

	if hours := calculateWorkingHours(6, 17, 5); hours >= 52 {
		fmt.Printf("Hours spent with coding is more than 52, hours %.2f\n", hours)
	}

	a := 15
	b := 88
	fmt.Printf("a=%d, b=%d\n", a, b)
	a, b = swapVariables(&a, &b)
	fmt.Printf("a=%d, b=%d\n", a, b)

	secondsFromDay := remainingSecondsFromDay()
	fmt.Println(secondsFromDay)

	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(filterNumbers(numbers, isEven))
	fmt.Println(filterNumbers(numbers, isOdd))

	checkDay(time.Now().Weekday())

	multiplicationTable(15)

	countFrom(2, 1)
	countFrom(2, 9)

	colorsWithShades := [3][]string{{"lime", "forest green", "olive"}, {"orange red", "red", "tomato"}, {"orchid", "violet"}}
	for row := range colorsWithShades {
		for column := range colorsWithShades[row] {
			fmt.Printf("%s ", colorsWithShades[row][column])
		}
		fmt.Println()
	}

	printMatrix()
}
