package main

import "flag"
import "fmt"

func main() {
    var flag_var int
    flag.IntVar(&flag_var, "flag_var", 0,  "help message for flag")
    flag.Parse()
    fmt.Println(flag_var)
}
