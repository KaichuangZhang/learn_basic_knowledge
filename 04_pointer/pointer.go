package main

import "fmt"


// 1. pointer in go, pointer can not change. (like ++, --, += ....)
// 2. pointer in go is easy.

func modify1(x int) {
    x = 100
}

func modify2(y *int) {
    *y = 100
}
func main() {

    // 1. basic 
    a := 10
    p_a := &a
    fmt.Printf("%T = %d(%p), %T = %d(%x)\n", a, a, &a,  p_a, *p_a, p_a)


    // 2. classic
    aa := 1
    modify1(aa)
    fmt.Println(aa)
    modify2(&aa)
    fmt.Println(aa)
}
