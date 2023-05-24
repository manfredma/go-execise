package main

import "fmt"

func main() {
	fmt.Println("hello, world")

	var x = []int{1, 2, 3, 4}

	var x2 = make([]int, 5, 10)
	copy(x2, x)

	fmt.Println("x=", x)
	fmt.Println("x2=", x2)

	var y = len(x)
	fmt.Println("len=", y)

	var z = x[:2]
	fmt.Println("z=", z)

	z = append(z, 3)

	fmt.Println("x=", z)
	fmt.Println("z=", z)
}
