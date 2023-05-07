package main

import (
    "fmt"
    "os"
    "log"
)


func main() {
    file, err := os.Open("api.go") // For read access.
    if err != nil {
        fmt.Println(err)
        log.Fatal(err)
    }
    fmt.Println(file)
}
