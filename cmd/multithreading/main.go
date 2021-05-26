package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	go f(i)
	// }

	// var c chan string = make(chan string)

	// go pinger(c)
	// go ponger(c)
	// go printer(c)

	c1 := make(chan string)
	c2 := make(chan string)
	c := make(chan string, 1)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()
	go func() {
		for {
			c <- "Buffer"
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			c <- "Buffer2"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case msg3 := <-c:
				fmt.Println("Buffer", msg3)
			case <-Sleep(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

func Sleep(sec time.Duration) <-chan time.Time {
	return time.After(sec)
}

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
