package main

import "fmt"

var dogsName = "Ted"

func input() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func degree() {
	fmt.Print("Enter a Fahrenheit: ")
	var F int
	fmt.Scanf("%f", F)

	C := (F - 32) * 5 / 9

	fmt.Println("Celsius: ", C)
}

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	var y string
	y = "Hello World!"
	fmt.Println(y)

	var a string
	a = "first"
	fmt.Println(a)
	a += "second"
	fmt.Println(a)

	var b string = "hello"
	var c string = "hello"
	fmt.Println(b == c)

	d := "World"
	fmt.Println(d)

	fmt.Println(dogsName)
	f()
	input()
	degree()
}

func f() {
	fmt.Println(dogsName)

	const x string = "Hello world!!!!!!"
	fmt.Println(x)

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)
}
