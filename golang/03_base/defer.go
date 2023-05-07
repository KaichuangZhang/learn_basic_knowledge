package main

import (
    "fmt"
)

// defer
// sourece release
func test_defer() {
    fmt.Println("start.")
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    fmt.Println("end.")
}


func main() {
    test_defer()
}
