package main

import "fmt"

func main() {
	fmt.Println("int: 1 + 1 = ", 1+1)
	fmt.Println("int: 1 - 1 = ", 1-1)
	fmt.Println("int: 4 / 2 = ", 4/2)
	fmt.Println("int: 5 * 5 = ", 5*5)
	fmt.Println("int: 4 % 3 = ", 4%3)
	fmt.Println(" \n ")

	fmt.Println("int: 1.1 + 1.2 = ", 1.1+1.2)
	fmt.Println("int: 1.1 - 1.2 = ", 1.1-1.2)
	fmt.Println("int: 4.1 / 2.2 = ", 4.1/2.2)
	fmt.Println("int: 5.1 * 5.2 = ", 5.1*5.2)
	fmt.Println(32132 * 42452)
	fmt.Println(" \n ")

	fmt.Println("string: \n Hello \t World")
	fmt.Println(`
		string
		with
		brackets
	`)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1]) // e = 101 byte
	fmt.Println("Hello " + "World")
	fmt.Println(" \n ")

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println(false || false || true)
	fmt.Println((true && false) || (false && true) || !(false && false))
}
