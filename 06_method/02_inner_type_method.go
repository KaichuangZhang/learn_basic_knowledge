package main

import (
    "fmt"
)

// the method of int
// error :
//func (i int) my_method() {
//    fmt.Println("this is my method of int")
//}
// the right way:
type myInt int
func (i myInt) my_method() {
    fmt.Println("this is my method of int")
}

func main() {
    var a myInt
    a = 10
    a.my_method()
}

