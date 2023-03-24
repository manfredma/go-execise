package main

import "fmt"

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("LL: 异常信息为:", err)
		}
	}()

	fmt.Printf("Multiply 2 * 5 * 6 = %d\n", MultiPly3Nums(2, 5, 6))
	// var i1 int = MultiPly3Nums(2, 5, 6)
	// fmt.Printf("MultiPly 2 * 5 * 6 = %d\n", i1)

	TestPanic()
}

func MultiPly3Nums(a int, b int, c int) int {
	// var product int = a * b * c
	// return product
	return a * b * c
}

func TestPanic() {

	panic("发生恐慌了")
}
