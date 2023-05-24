package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Begin~")
	ifEx()
	forEx()
	switchEx()
}

func ifEx() {
	x := rand.Intn(10)
	if x == 0 {
		fmt.Println("That's too low")
	} else if x > 5 {
		fmt.Println("That's too big: ", x)
	} else {
		fmt.Println("That's a good number: ", x)
	}
}

func forEx() {
	var s = []string{"ABC", "DEF", "GHI", "apple_π！"}
	for _, v := range s {
		fmt.Println("s=", v)
		for i, c := range v {
			fmt.Println("i=", i, ", c=", c, ", str=", string(c))
		}
	}

	var m = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	fmt.Println("+__+ 1")
	for k, v := range m {
		fmt.Println("key=", k, ", value=", v)
	}
	fmt.Println("+__+ 2")
	for k, v := range m {
		fmt.Println("key=", k, ", value=", v)
	}
}

func switchEx() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short world!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")

		}
	}
}
