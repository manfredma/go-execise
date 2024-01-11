package main

import "fmt"

func main() {
	var intVar = 8
	fmt.Println("intVar: ", intVar)
	fmt.Println("intPointer: ", &intVar)

	var structVar = struct {
		age     int8
		name    string
		address struct{}
	}{
		age:  18,
		name: "manfred",
	}
	fmt.Println("structVar: ", structVar)
	fmt.Println("structPointer: ", &structVar)
	var structPointer = &structVar
	fmt.Println("structPointerPointer: ", &structPointer)
}
