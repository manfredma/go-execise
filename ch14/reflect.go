package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f Foo
	ft := reflect.TypeOf(f)
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(), curField.Tag.Get("myTag"))
	}

	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	ivv.SetInt(20)
	fmt.Println(i)
}

type Foo struct {
	A int    `myTag:"value"`
	B string `myTag:"value2"`
}
