package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x float64
	y float64
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.area())
	fmt.Println(c)

	p := Person{Name: "Vasya"}

	a := Android{Model: "456"}

	fmt.Println(totalTalk(&p, &a))
}

type Person struct {
	Name string
}

func (p *Person) Talk() string {
	return p.Name
}

type Android struct {
	Model string
}

func (a *Android) Talk() string {
	return a.Model
}

type WhoCanTalk interface {
	Talk() string
}

func totalTalk(objects ...WhoCanTalk) string {
	var talk string
	for _, v := range objects {
		talk += v.Talk()
	}
	return talk
}
