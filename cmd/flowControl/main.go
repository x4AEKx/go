package main

import "fmt"

func main() {
	divisibleByThree()
	fizzBuzz()
	switchFunc()
}

func divisibleByThree() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz ", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz ", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz ", i)
		}
	}
}

func switchFunc() {
	for i := 0; i < 7; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}

}
