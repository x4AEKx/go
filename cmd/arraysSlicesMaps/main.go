package main

import "fmt"

func main() {
	arrays()
	slice()
	maps()
	smallestElement()
}

func arrays() {
	x := [5]float64{100, 54, 12, 35, 23}

	var total float64 = 0
	for _, v := range x {
		total += v
	}

	fmt.Println(total / float64(len(x)))
	fmt.Println(x)

	a := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(a[2:5])
}

func slice() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice5[4])

	slice6 := make([]int, 3, 9)
	fmt.Println(len(slice6))
}

func maps() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	delete(x, "key")
	fmt.Println(x)

	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

func smallestElement() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	y := x[0]

	for _, v := range x {
		if y > v {
			y = v
		} else if y == 0 {
			y = v
		}
	}

	fmt.Println(y)
}
