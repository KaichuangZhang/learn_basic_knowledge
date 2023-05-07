package main

import (
    "fmt"
)

// lambda function
// func (paramters) (return) {
//      function body
// }

// 闭包 = function + 外层变量
func test_f(name string) func() {
    return func() {
        fmt.Println("hello! bibao", name)
    }
}
func main() {
    // lambda function
    sayhello := func(){
        fmt.Println("hello.")
    }
    sayhello()
    func() {
        fmt.Println("hello.")
    }()
    
    // 闭包
    sayhi := test_f("zkc")
    sayhi()
}

