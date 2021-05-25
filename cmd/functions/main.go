package main

import "fmt"

func main() {
	defer fmt.Println(recursionFactorial(5))

	xs := []float64{98, 93, 77, 82, 83}

	fmt.Println(average(xs))
	fmt.Println(add(1, 2, 3))

	x := 0
	closure := func() int {
		x++
		return x
	}
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())
	fmt.Println(closure())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	defer fmt.Println(sum([]int{15, 12, 400}))

	defer fmt.Println(EvenOdd(2))

	defer fmt.Println("Fibonacy: ", fibonacy(555))

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func average(slice []float64) float64 {
	total := 0.0
	for _, v := range slice {
		total += v
	}
	return total / float64(len(slice))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func sum(x []int) int {
	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum
}

func EvenOdd(x int) (int, bool) {
	newVar := x / 2
	if x%2 == 0 {
		return newVar, true
	} else {
		return newVar, false
	}
}

func largestValue(args ...int) int {
	var x int = 0
	for _, v := range args {
		if x < v {
			x = v
		} else if x == 0 {
			x = v
		}
	}
	return x
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func recursionFactorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * recursionFactorial(x-1)
}

func fibonacy(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibonacy(n-1) + fibonacy(n-2)
}
