package main

import "fmt"

// 1. struct define
type person struct {
    name string
    age int
    city string
}

func main() {
    // declear
    var p1 person
    p1.name = "zkc"
    p1.age = 30
    p1.city = "bj"
    fmt.Printf("%v\n", p1)

    // lambda struct
    var p2 struct{
        first_name, last_name string
    }
    p2.first_name = "zhang"
    p2.last_name = "kaichuang"
    fmt.Println(p2)

    // struct pointer
    var p3 = new(person)
    (*p3).name = "zkc2"
    (*p3).age = 31
    (*p3).city = "beijing"
    fmt.Println(*p3)
    p3.name = "zkc3"
    p3.age = 32
    p3.city = "beijing2"
    fmt.Println(*p3)


    // init
    // 1. key: value
    // 2. value list
    p4 :=  person{name: "zhangkaichuang", age: 33, city: "bjjj"}
    fmt.Println(p4)
    var p5 person
    p5 = person{name: "zhangkaichuang33", age: 35, city: "bbbbjjj"}
    fmt.Println(p5)
    p6 := person{"zhangkaichuang", 36, "bbbbjjj"}
    fmt.Println(p6)


    // struct new function
    // self define new function

}   
