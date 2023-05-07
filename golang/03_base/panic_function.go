package main

import "fmt"

// panic
func f1() {
    fmt.Println("f1")
}

func f2() {
    defer func() {
        err := recover()
        if err != nil {
            fmt.Println("f2 error")
        }
    }()
    panic("error in f2")
}

func f3() {
    fmt.Println("f3")
}

// it is different if recover have or not
func main() {
    f1()
    f2()
    f3()
}
