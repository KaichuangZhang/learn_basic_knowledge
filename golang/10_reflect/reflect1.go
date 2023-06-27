package main

import (
	"fmt"
	"reflect"
)

func returnType(x interface{}) {
	// 1. 类型断言
	// 2. reflect.Typeof
	t := reflect.TypeOf(x)
	fmt.Println(t, t.Name(), t.Kind())
	fmt.Printf("%T\n", t)
}

func main() {
	var a float32 = 1.234
	returnType(a)

	var b int8 = 10
	returnType(b)

}
