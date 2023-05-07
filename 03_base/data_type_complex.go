package main

import (
    "fmt"
)
func main(){
    var array [10]int // 可以把 [10]int看成一种类型
    var array2 [10]int
    if array == array2 {
        fmt.Println(array, array2)
        fmt.Printf("%T %T\n", array, array2)
    }
    // init
    var CityArray = [4]string{"Bei", "SHang", "Gz", "SZ"}
    fmt.Println(CityArray)
    var CArray = [...]bool{false, true}
    fmt.Println(CArray)
    fmt.Printf("%d\n", len(CArray))
    CArray[0] = true
    fmt.Println(CArray)
    var BArray = [...]bool{1:false, 10:true}
    fmt.Println(BArray)
    fmt.Printf("%d\n", len(BArray))
    for index, value := range BArray {
        fmt.Printf("%d:%t\n", index, value)
    }
    // [...][10] is ok, but [...][...] is not ok
    var B2Array = [...][11]bool{
        {1:false, 10:true},
        {1:false, 10:true},
    }
    fmt.Println(B2Array)
    for index1, B2 := range B2Array {
        for index2, BB2 := range B2 {
            fmt.Printf("[%d][%d] %t",index1, index2, BB2)
            
        }
        fmt.Println()
    }
}
