package main

import (
    "fmt"
    "sort"
    "math/rand"
)

func main() {
    // declare
    var map0 map[string]int
    fmt.Println(map0 == nil)
    
    // define
    map0 = make(map[string]int, 8)
    fmt.Println(map0)
    map0["one"] = 1
    map0["two"] = 2
    fmt.Println(map0)

    // declare & define
    map1 := map[int]bool {
        1: true,
        2: false,
        3: true,
    }
    fmt.Println(map1)
    
    // error example
    // var map2 map[string]int
    // map2["key"] = 1 // will error
    
    // if key is exist or not
    var map3 = make(map[string]int, 10)
    map3["key1"] = 100
    map3["key2"] = 200
    value, ok := map3["key3"]
    fmt.Println(value, ok)
    value, ok = map3["key2"]
    fmt.Println(value, ok)
    
    // iterator
    for key, value := range map3 {
        fmt.Println(key, value)
    }
    for key := range map3 {
        fmt.Println(key)
    }
    for _, value := range map3 {
        fmt.Println(value)
    }

    // delete
    delete(map3, "key1")
    fmt.Println(map3)
    
    // exmpale : iterate by sorted key
    stu := make(map[string]int, 50)
    for i := 0; i < 50; i ++ {
        key := fmt.Sprintf("student%2d", i)
        value := rand.Intn(100)
        stu[key] = value
    }
    fmt.Println(stu)
    keys := make([]string, 0, 50)
    for key := range stu {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    for _, key := range keys {
        fmt.Println(key, stu[key])
    }

    // slice of map
    // map of slice

    // example: count the char 
    fmt.Println("--------")
    str := "122333444455555666666777777788888888"
    fmt.Println(len(str))
    map_str_count := make(map[rune]int, 26)
    for _, ch := range str {
        _, ok := map_str_count[ch]
        //fmt.Println(ok)
        if ok {
            map_str_count[ch] += 1
        } else {
            map_str_count[ch] = 1
        }
    }
    fmt.Println(len(map_str_count))
    fmt.Println(map_str_count)
}

