package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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
	for _, num := range numbers {
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

func findMatchingIndexes(input int, numbers []int) []int {
	var matches []int
	for i, num := range numbers {
		if isSameDigit(num, input) {
			matches = append(matches, i)
		}
	}
	return matches
}

func isSameDigit(num int, input int) bool {
	for num > 0 {
		digit := num % 10
		if digit == input {
			return true
		}
		num = num / 10
	}
	return false
}

func findUniqueNumbers(nums []int) []int {
	numbers := make(map[int]int)
	var uniques []int

	for _, num := range nums {
		if numbers[num] == 0 {
			uniques = append(uniques, num)
			numbers[num] = 1
		} else {
			numbers[num]++
		}
	}

	return uniques
}

func isAnagram(s1 string, s2 string) bool {
	first, second := []byte(s1), []byte(s2)
	sort.Slice(first, func(i, j int) bool { return first[i] < first[j] })
	sort.Slice(second, func(i, j int) bool { return second[i] < second[j] })
	return string(first) == string(second)
}

func generatePalindrome(s string) string {
	var builder strings.Builder
	builder.WriteString(s)

	bytes := []byte(s)
	for i := len(bytes) - 1; i >= 0; i-- {
		builder.WriteByte(bytes[i])
	}

	return builder.String()
}

func fixUrl(url *string) string {
	tokens := strings.Split(*url, "/")
	tokens[0] += ":"
	tokens[len(tokens)-1] = strings.Replace(tokens[len(tokens)-1], "bots", "odds", 1)
	*url = strings.Join(tokens, "/")
	return *url
}

func reverse(s *string) string {
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func sum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func isArmstrongNumber(n int) bool {
	var digits []int

	tmp := n
	for tmp != 0 {
		digit := tmp % 10
		tmp /= 10
		digits = append(digits, digit)
	}

	size := len(digits)
	var digitsOnPow []float64
	for _, digit := range digits {
		digitsOnPow = append(digitsOnPow, math.Pow(float64(digit), float64(size)))
	}

	return sum(digitsOnPow) == float64(n)
}

func divideByTen(n float64) (float64, error) {
	if n == 0 {
		return -1, errors.New("math: can't divide by zero")
	}
	return n / 10, nil
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

	animals := [15]string{"koal", "pand", "zebr", "anacond", "bo", "chinchill", "cobr", "gorill", "hyen", "hydr", "iguan",
		"impal", "pum", "tarantul", "pirahn"}
	for i, animal := range animals {
		animal += "a"
		animals[i] = animal
	}
	fmt.Println(animals)

	orders := [3]string{"first", "second", "third"}
	tmp := orders[0]
	orders[0] = orders[2]
	orders[2] = tmp
	fmt.Println(orders)

	fmt.Println(findMatchingIndexes(1, []int{1, 11, 34, 52, 61}))
	fmt.Println(findMatchingIndexes(9, []int{1, 11, 34, 52, 61}))
	fmt.Println(findMatchingIndexes(5, []int{1, 11, 34, 52, 61}))

	fmt.Println(findUniqueNumbers([]int{1, 11, 34, 11, 52, 61, 1, 34}))

	fmt.Println(isAnagram("dog", "god"))
	fmt.Println(isAnagram("pear", "apple"))

	fmt.Println(generatePalindrome("santaclaus"))

	wrongText := "In a dishwasher far far away"
	wrongText = strings.Replace(wrongText, "dishwasher", "galaxy", 1)
	fmt.Println(wrongText)

	wrongRedditUrl := "https//www.reddit.com/r/nevertellmethebots"
	fixUrl(&wrongRedditUrl)
	fmt.Println(wrongRedditUrl)

	wrongQuote := "Hofstadter's Law: It you expect, even when you take into account Hofstadter's Law."
	after := strings.SplitAfter(wrongQuote, "It")
	if len(after) > 1 {
		after[0] += " always takes longer than"
	}
	wrongQuote = strings.Join(after, "")
	fmt.Println(wrongQuote)

	text := ".eslaf eb t'ndluow ecnetnes siht ,dehctiws erew eslaf dna eurt fo sgninaem eht fI"
	fmt.Println(reverse(&text))

	fmt.Println(isArmstrongNumber(1634))
	fmt.Println(isArmstrongNumber(371))
	fmt.Println(isArmstrongNumber(153))
	fmt.Println(isArmstrongNumber(4354432422))

	divided, err := divideByTen(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(divided)
	}
}
