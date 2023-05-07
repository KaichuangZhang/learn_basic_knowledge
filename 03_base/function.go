package main

import (
	"fmt"
)

// define
// func func_name(parameters)(reture) {
//      function body
// }

// 1.
func sayhello() {
	fmt.Println("hello main")
}

func sayhelloto(name string) {
	fmt.Println("hello,", name)
}

// 2.
func sum(x int, y int) int {
	return x + y
}

func sum2(x int, y int) (sum int) {
	sum = x + y
	return
}

// 3. varible parameter
func sum3(x ...int) (sum int) {
	fmt.Printf("%T\n", x)
	for _, x_v := range x {
		sum += x_v
	}
	return
}

// 4. have no default paramter
// type simple
func sum4(x, y int) (sum int) {
	sum = x + y
	return
}

// 5.  mul return value
func calc(x, y int) (sum, mu int) {
	sum = x + y
	mu = x - y
	return
}


func main() {
	sayhello()
	sayhelloto("main")
	fmt.Println(sum(1, 2))
	fmt.Println(sum3(1, 2))
	fmt.Println(sum3())
	fmt.Println(calc(10, 2))
}
