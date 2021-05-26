package main

import (
	"fmt"

	"github.com/x4aekx/go/pkg/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
