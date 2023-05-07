package main

import (
    "fmt"
    "sort"
)

// type - slice
//        []T
func main() {
    // 1.
    var slice1 []int
    var slice2 []string
    var slice3 = []bool{false, true}
    fmt.Println(slice1)
    fmt.Println(slice2)
    fmt.Println(slice3)


    // 2. from array
    array := [5]int{1, 2, 3, 4, 6}
    slice4 := array[1:4]
    fmt.Printf("%T %T\n", array, slice4)
    fmt.Println(slice4)

    // 3. from slice
    slice5 := slice4[1:2]
    fmt.Printf("%T %T\n", slice4, slice5)
    fmt.Println(slice5)

    // 4. make
    slice6 := make([]int, 5, 10)
    fmt.Printf("%T\n", slice6)
    fmt.Println(slice6)

    // ------- fucntion of slice --------
    // len
    slice6_len := len(slice6)
    fmt.Printf("%d", slice6_len)
    fmt.Println(slice6_len)
    // cap


    // nil slice
    var slice7 []int
    fmt.Println(slice7, len(slice7), cap(slice7))
    if slice7 == nil {
        fmt.Println("slice7 is nil.")
    }
    
    var slice8 = []int{}
    if slice8 == nil {
        fmt.Println("slice8 is nil")
    } else {
        fmt.Println("slice8 is not nil")
    }

    // = copy
    slice9 := make([]int, 3)
    slice10 := slice9
    fmt.Println(slice9, slice10)
    slice9[1] = 100
    fmt.Println(slice9, slice10)


    // iterator
    slice11 := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(slice11); i ++ {
        fmt.Println(i, slice11[i])
    }
    fmt.Println()
    for i, value := range slice11 {
        fmt.Println(i, value)
    }
    fmt.Println()
    
    // append
    slice11 = append(slice11, 6)
    for i, value := range slice11 {
        fmt.Println(i, value)
    }
    fmt.Println()
    slice11 = append(slice11, 7, 8, 9, 10, 11)
    for i, value := range slice11 {
        fmt.Println(i, value)
    }
    slice11 = append(slice11, slice10...)
    for i, value := range slice11 {
        fmt.Println(i, value)
    }


    // copy - deep copy
    slice_0 := []int{1, 2, 3, 4, 5}
    slice_1 := make([]int, 5, 5)
    copy(slice_1, slice_0)
    fmt.Println(slice_0, slice_1)
    slice_0[1] = 100
    fmt.Println(slice_0, slice_1)

    // delete
    slice_2 := []int{1, 2, 3, 4, 5}
    slice_2 = append(slice_2[0:3], slice_2[4:]...)
    fmt.Println(slice_2)
    
    sort.Ints(slice_1)
    fmt.Println(slice_1)
}
