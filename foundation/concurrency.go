package main

import (
	"fmt"
	"runtime"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func sumNumbers(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// always close channels in producers
	close(c)
}

func fibonacciTwoChannels(quit, c chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	go say("world")
	say("hello")

	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sumNumbers(a[:len(a)/2], c)
	go sumNumbers(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	ch := make(chan int, 10)
	go fibonacci(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}

	channel := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-channel)
		}
		quit <- 0
	}()
	fibonacciTwoChannels(quit, channel)

	channel2 := make(chan int)
	timeout := make(chan bool)
	go func() {
		for {
			select {
			case v := <-channel2:
				fmt.Println(v)
			case <-time.After(time.Second * 3):
				fmt.Println("timeout")
				timeout <- true
				break
			}
		}
	}()
	<-timeout

	fmt.Printf("Number of CPUs is %d", runtime.NumCPU())
}
