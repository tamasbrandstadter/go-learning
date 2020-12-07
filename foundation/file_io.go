package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type kv struct {
	Key   string
	Value int
}

func printFileContent(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return errors.New("can't open file " + fileName)
	} else {
		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()

		buf := make([]byte, 1024)
		for {
			n, err := file.Read(buf)

			if n > 0 {
				fmt.Println(string(buf[:n]))
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("read %d bytes: %v", n, err)
				break
			}
		}
		return nil
	}
}

func countLinesInFile(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, errors.New("can't open file " + fileName)
	} else {
		defer func() {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}()

		buf := make([]byte, 32*1024)
		count := 0
		separator := []byte{'\n'}
		for {
			c, err := file.Read(buf)
			count += bytes.Count(buf[:c], separator)

			switch {
			case err == io.EOF:
				return count, nil

			case err != nil:
				return count, err
			}
		}
	}
}

func writeContentToFile(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	if _, err := file.WriteString(content); err != nil {
		log.Println(err)
	}
}

func copyFileContent(fileFrom string, fileTo string) bool {
	content, err := ioutil.ReadFile(fileFrom)
	if err != nil {
		return false
	}
	writeContentToFile(fileTo, string(content))
	return true
}

func fiveMostCommonLotteryNumbers(fileName string) ([5]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return [5]string{}, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	frequency := make(map[string]int)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\r')
		if err != nil || line == "" {
			break
		}

		tokens := strings.Split(line, ";")
		size := len(tokens)

		frequency[strings.TrimSuffix(tokens[size-5], "\r")]++
		frequency[strings.TrimSuffix(tokens[size-4], "\r")]++
		frequency[strings.TrimSuffix(tokens[size-3], "\r")]++
		frequency[strings.TrimSuffix(tokens[size-2], "\r")]++
		frequency[strings.TrimSuffix(tokens[size-1], "\r")]++
	}

	var numbers []kv
	for number, count := range frequency {
		numbers = append(numbers, kv{number, count})
	}

	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].Value > numbers[j].Value
	})

	return [5]string{numbers[0].Key, numbers[1].Key, numbers[2].Key, numbers[3].Key, numbers[4].Key}, nil
}

func main() {
	fileName := "my-file.txt"

	if err := printFileContent(fileName); err != nil {
		fmt.Println(err)
	}

	if lineCount, err := countLinesInFile(fileName); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Line count %d\n", lineCount)
	}

	writeContentToFile("not-existing.txt", "apple")
	writeContentToFile(fileName, "pear")

	copyFileContent(fileName, "test.txt")

	if numbers, err := fiveMostCommonLotteryNumbers("lottery.csv"); err == nil {
		fmt.Println(numbers)
	}
}
