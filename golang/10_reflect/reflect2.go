package main

import (
	"fmt"
	"reflect"
)

func returnValue(x interface{}) {
	// 1. 类型断言
	// 2. reflect.Typeof
	v := reflect.ValueOf(x) // type is reflect.Value, value is value.
	fmt.Println(v)
	fmt.Printf("%T\n", v)
	switch v.Elem().Kind() {
	case reflect.Float32:
		//v2 := v.Float()
		v.Elem().SetFloat(1.00)
		fmt.Printf("%#v, %T\n", v.Elem(), v)
		break
	case reflect.Int:
		v.Elem().SetInt(1)
		fmt.Printf("%#v, %T\n", v.Elem(), v)
		break
	default:
		break
	}
}

func main() {
	var a float32 = 1.234
	returnValue(&a)

	var b int8 = 10
	returnValue(&b)

}
