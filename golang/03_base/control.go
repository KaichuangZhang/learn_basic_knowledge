package main

import (
    "fmt"
)


func main(){
    // 分支
    // if else
    if false {
        fmt.Println("1")
    } else if false {
        fmt.Println("2")
    } else {
        fmt.Println("3")
    }
    // 循环
    fmt.Print("=======\n")
    // for
    for i := 1; i <= 10; i ++ {
        fmt.Println(i)
    }
    // while
    fmt.Print("=======\n")
    i := 1
    for i < 10 {
        fmt.Println(i)
        i += 1
    }
    fmt.Print("=======\n")
    i = 1
    for i < 10 {
        fmt.Println(i)
        i += 1
        if i >= 5 {
            break
        }
    }
    // for range
    // switch
    // switch case 1, 2, 3, 4
    // switch 变量
}

