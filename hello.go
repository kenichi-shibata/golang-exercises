// https://www.youtube.com/watch?v=C8LgvuEBraI
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello, 世界")

	// variables
	var y int = 2
	x := 3

	fmt.Println(y)
	fmt.Println(x)

	// array variables
	var arr [3]int
	arr2 := [3]int{0, 12, 2}

	fmt.Println(arr)
	fmt.Println(arr2)

	// if
	z := 3
	if z > 2 {
		fmt.Println("z is greater than 2")
	}

	// if else with (no truthies)
	a := ""

	if a != "" {
		fmt.Println("a exists")
	} else {
		fmt.Println("a does not exists!")
	}

	// for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}

	fmt.Println()
	// while loop
	h := 0

	for h < 5 {
		fmt.Print(h)
		h++
	}

	fmt.Println()
	// loop on array
	arr3 := []string{"a", "b", "c", "d"}

	for index, value := range arr3 {
		fmt.Println("index", index, "value", value)
	}

	// map
	m := make(map[string]int)

	m["apple"] = 1
	m["ball"] = 2
	fmt.Print(m)
	fmt.Print(m["ball"])

	fmt.Println()
	// map looping
	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	// call function
	result := sum(x, y)
	fmt.Println(result)

	sqrtResult, err := sqrt(-1)
	if err != nil {
		fmt.Println("sqrtErr", err)
	} else {
		fmt.Println(sqrtResult)
	}

	sqrtResult, err = sqrt(10)
	if err != nil {
		fmt.Println("sqrtErr", err)
	} else {
		fmt.Println(sqrtResult)
	}

	// struct
	p := person{name: "Jake", age: 15}
	fmt.Println(p)
	fmt.Println(p.name)

	// pointers
	g := 7
	fmt.Println(g)
	increment(&g)
	fmt.Println(g)

}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

// pointers
func increment(x *int) int {
	(*x)++
	return *x
}

// struct
type person struct {
	name string
	age  int
}
