package main

import "fmt"
func main() {
    // 1.算术运算符
    a := 10
    b := 20
    c := a + b
    d := a - b
    e := a * b
    f := a / b
    g := a % b
    fmt.Println(a, b, c, d, e, f, g)
    a ++
    b --
    fmt.Println(a, b)
    // 2.逻辑运算符
    // == != >= <= > <
    // 3.逻辑运算符
    // && || !
    // 4.位运算
    // & | ^ << >>
    // 5. 赋值运算
    aa := 4
    aa += 2
    fmt.Println(aa)
}
