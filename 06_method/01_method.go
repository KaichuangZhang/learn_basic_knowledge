package main

import (
    "fmt"
)

// method is diffent from function
// method = function + reciver

// 1. struct define
type Person struct {
    name string
    age int
    city string
}

func newPerson(name string, age int, city string) *Person {
    return &Person {
        name: name,
        age: age,
        city: city,
    }
}
// the diffence of *Person and Person
func (p Person) change_name1(new_name string) {
    p.name = new_name
}
func (p *Person) change_name2(new_name string) {
    p.name = new_name
}
func main() {
    p1 := newPerson("zkc", 30, "bj")
    fmt.Println(*p1)
    p1.change_name1("zhangkaichuang") // change_name1 will not change p1
    fmt.Println(*p1)
    p1.change_name2("zhangkaichuang") // change_name2 will change p1
    fmt.Println(*p1)
}
