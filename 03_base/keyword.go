package main

import "fmt"
// keyword
// select
// interface
// defer
// chan
// fallthrough
// range
// var
// 保留字
// var
// 1. var name type
// for example:
// var name string
// var age int
// var isOk bool
// 2. var (a string b int c bool)
// iota (only can be used in const)
const (
    monday = iota
    t = iota
    
)
const pi = 3.4
func main() {
    var name string
    name = "zhangkaichuang"
    var age = 33
    isOk := true
    fmt.Println(name, age, isOk, pi)
}
