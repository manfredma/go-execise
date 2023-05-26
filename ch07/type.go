package main

import "fmt"

func main() {
	p := Person{
		FirstName: "Fred",
		LastName:  "Fred-son",
		Age:       12,
	}
	fmt.Println(p.String())

	//
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(12))
	fmt.Println(it.Contains(2))
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (receiver Person) String() string {
	return fmt.Sprintf("%s %s, age %d", receiver.FirstName, receiver.LastName, receiver.Age)
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	// 值是无法判空的，只有指针才可以！
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}

}
