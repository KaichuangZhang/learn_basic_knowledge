package main

import "fmt"

// 1. self define type
type myInt int

// 2. other name of type
type myInt2 = int
func main() {
    // 1.
    var a myInt
    a = 10
    fmt.Printf("%T %d\n", a, a)
    // 2.
    var aa myInt2
    aa = 10
    fmt.Printf("%T %d\n", aa, aa)
}
