package main

import "fmt"

type dog struct {}
func (d *dog) call() {
    fmt.Println("wang wang ...")
}

type cat struct {}
func (c *cat) call() {
    fmt.Println("miao miao ...")
}

type person struct {
    name string
}
func (p *person) call() {
    fmt.Println("wawa wawa ...")
}

// type interface_name interface {
//      method_name(parameter_list) return_list
// }
type animal interface {
    call()
}

func let_call(an animal) {
    an.call()
}


// null interface
// and this type is not nessaray
// 1. as the parameter of function
// 2. map's value - it's value type is all
// interface = type + value
type null_type interface {
    // no method require
}
func main() {
    // 1
    d := dog{}
    c := cat{}
    let_call(&d)
    let_call(&c)

    p := person {
        name: "zhangkaichuang",
    }
    let_call(&p)


    // 2
    var call_type animal
    call_type = &cat{}
    call_type.call()
}
