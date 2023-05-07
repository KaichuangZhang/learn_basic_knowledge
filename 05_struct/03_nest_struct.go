package main

import "fmt"

// lambda struct
type Address struct {
    // province string
    // id int
    string
    int
}

type Address1 struct {
    Province string
    Id int
}

type Email struct {
    Addr string
    Id int
}

// nest struct
// lambda struct in nest struct
type Person struct {
    Name string
    Age int
    Address Address1
    Email Email
}


func main() {
    var p = Person {
        Name: "zkc",
        Age: 30,
        Address: Address1 {
            Province: "henan",
            Id: 123,
        },
        Email: Email {
            Addr : "zkccst@gmail.com",
            Id: 22,
        },
    }
    fmt.Println(p)
    // fmt.Println(p.Id) this is not right

    fmt.Println(p.Address.Id)
    fmt.Println(p.Email.Id)
}
