package main

import "fmt"

func main() {
	var a1 = [5]int{1, 2, 3, 4, 5}
	// 数组赋值是值拷贝！
	var a2 = a1
	a2[0] = 99
	fmt.Println(a1, a2)
	fmt.Println("1> ----------------------------a1[0]==a2[0]? ", a1[0] == a2[0]) // false

	// 切片
	var sa = a1[:]
	sa[1] = 88
	fmt.Println(a1, sa)
	fmt.Println("2> ----------------------------a1[1]==sa[1]? ", a1[1] == sa[1]) // true

	var s1 = []int{1, 2, 3, 4, 5}
	// slice赋值是引用"值"拷贝！
	var s2 = s1
	s2[0] = 99
	fmt.Println(s1, s2)
	fmt.Println("3> ----------------------------s1[0]==s2[0]? ", s1[0] == s2[0]) // true

	var s3 = s2[1:4]
	s3[0] = 88
	fmt.Println(s1, s3)
	fmt.Println("4> ----------------------------s1[1]==s3[0]? ", s1[1] == s3[0]) // true

	var s4 = append(s2, 6)
	s4[2] = 77
	fmt.Println(s1, s4)
	fmt.Println("5> ----------------------------s1[2]==s4[2]? ", s1[2] == s4[2]) // false

	var oldLen int
	oldLen = len(s1)
	fmt.Println(s1)
	appendInt(s1, 6)
	fmt.Println(s1)
	fmt.Println("6> ----------------------------len(s1)==oldLen+1?", len(s1) == oldLen+1) // false

	oldLen = len(s1)
	fmt.Println(s1, s2)
	appendIntPtr(&s2, 6)
	fmt.Println(s1, s2)
	fmt.Println("7> ----------------------------len(s1)==oldLen+1?", len(s1) == oldLen+1) // false

	oldLen = len(s1)
	fmt.Println(s1)
	appendIntPtr(&s1, 6)
	fmt.Println(s1)
	fmt.Println("8> ----------------------------len(s1)==oldLen+1?", len(s1) == oldLen+1) // true
}
func appendInt(s []int, elems ...int) {
	s = append(s, elems...)
}
func appendIntPtr(ps *[]int, elems ...int) {
	*ps = append(*ps, elems...)
}
