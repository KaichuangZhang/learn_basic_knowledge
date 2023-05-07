package main

import "fmt"

// the difference of new and make
// 1,
// 2,
// 3,
func main() {
    var p_a *int = new(int)
    // this is not right:
    // var p_a *int
    // *p_a = 10 
    *p_a = 10
    fmt.Println(*p_a)


    // map
    var p_b map[string]int
    p_b = make(map[string]int, 10)
    p_b["1"] = 100
    fmt.Println(p_b)
}
