/**
 * codehackfox@gmail.com
 * 2016/4/14
 * 选择排序
 */

package main

import "fmt"

func main() {
    arr := [10]int {4,6,2,3,0,12,23,45,13,9}
    fmt.Println(arr)
    arr_len  := len(arr)

    for i:=0;i<arr_len;i++{
        min := i
        flag :=true
        for j:=i;j<arr_len;j++{
            if arr[min]>arr[j]{
                flag=false
                min = j
            }
        }
        if flag {
            break
        }
        arr[i],arr[min] = arr[min],arr[i]
        fmt.Println(arr)
    }
}
