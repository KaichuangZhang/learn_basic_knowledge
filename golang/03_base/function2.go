package main

import (
    "fmt"
)

// global var
var num int = 100

func testgloble() {
    fmt.Println(num) // num is a globle var
}

// var fun
func testvarfun() {
    fmt.Println("test func var")
}

// func as paramters

func add(x, y int) int {
    return x + y
}

func sub(x, y int) int {
    return x - y
}

func calc(x, y int, op func(int, int) int) int {
    return op(x, y)
}
func main() {
    // 
    testgloble()
    num += 1
    testgloble()



    // test var func
    funct := testvarfun
    funct()
    
    

    // test func as p
    fmt.Println(calc(10, 2, add))
    fmt.Println(calc(10, 2, sub))
}
