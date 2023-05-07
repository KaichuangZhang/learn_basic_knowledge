package main

import (
    "fmt"
    "math"
)

func main() {
    // int - uint8 uint16 uint32 uint64 int8 int16 int32 int64 - uint(32 system - uint32) int uinptr (uint poniter)
    n := 10
    fmt.Printf("%b\n", n)
    m := 075
    fmt.Printf("%d\n", m)
    fmt.Printf("%o\n", m)
    fmt.Printf("%x\n", m)
    // float - float32 flot64
    fmt.Println(math.MaxFloat32, math.MaxFloat64)
    // complex - complex32 complex64
    // bool - true false
    var flag bool
    fmt.Println(flag)
    flag = true
    fmt.Println(flag)
    // string 
    s1 := "hello go"
    fmt.Println(s1)
    s2 := `
        uou
        ewe
    `
    fmt.Println(s2)
    fmt.Println(s1 + s2)
    fmt.Sprintf("%s%s", s1, s2)
    // strings.Split()  - input and output
    // strings.Contains()
    // strings.HasPref()
    // strings.HasSuffix()
    // Index
    // LastIndex
    // Join
    fmt.Println("len = ", len(s1))
    fmt.Println("len = ", len(s2))
    
    // 字符
    // 转义字符
    path := "c:\\code\\github"
    fmt.Println(path)
    // byte - uint ASCII
    var c1 byte = 'c'
    // rune - int32
    var c2 rune = 'c'
    fmt.Printf("c1:%T, c2:%T\n", c1, c2)
    c3 := "1234 你妈妈们"
    for i:=0; i < len(c3); i += 1 {
        fmt.Printf("%c ", c3[i])
    }
    fmt.Println()
    for i, r := range c3 {
        //fmt.Printf("%c ", r)
        fmt.Printf("%d- %c\n", i, r)
        //fmt.Println(i, r)
    }
}
