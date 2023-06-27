package main

import (
	"fmt"
	"reflect"
)

type Dog struct {
	name string `json:"name"`
	id   int8   `json:"id"`
}

func (d Dog) printinfo() string {
	result := fmt.Sprintf("my name is %s, and my id is %d. i am a dog.\n", d.name, d.id)
	return result
}

func getMethod(x interface{}) {
	dtype := reflect.TypeOf(x)
	dvalue := reflect.ValueOf(x)
	fmt.Printf("%#v\n", dtype)
	fmt.Printf("%#v\n", dvalue)
	// method
	fmt.Println(dtype.NumMethod())
	fmt.Printf("the number of method of dog is %d.", dvalue.NumMethod())
	for i := 0; i < dtype.NumMethod(); i += 1 {
		// get method by id
		method := dtype.Method(i)
		fmt.Println(method.Name, method.Type)
	}
}

func getField(x interface{}) {
	dtype := reflect.TypeOf(x)
	// field
	for i := 0; i < dtype.NumField(); i += 1 {

		// get field by id
		field := dtype.Field(i)
		fmt.Println(field.Name, field.Tag, field.Tag.Get("json"))

		// get field by name
		field2, ok := dtype.FieldByName("name")
		if ok {
			fmt.Println(field2.Name, field2.Tag, field2.Tag.Get("json"))
		}
	}
}

func main() {

	d := Dog{
		name: "zkc",
		id:   1,
	}
	d.printinfo()
	getField(d)
	getMethod(d) // TODO: get method is not ok.

}
