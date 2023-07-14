// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")

	var y int = 2
	x := 3

	fmt.Println(y)
	fmt.Println(x)

	var arr [3]int
	arr2 := [3]int{0, 12, 2}

	fmt.Println(arr)
	fmt.Println(arr2)

	z := 3
	if z > 2 {
		fmt.Println("z is greater than 2")
	}

	a := ""

	if a != "" {
		fmt.Println("a exists")
	} else {
		fmt.Println("a does not exists!")
	}


	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

