package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 29 * 10

func main() {
	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)

	// 编译错误，常量不能重新赋值！
	// x = x+1
	// y = "bye"

}
