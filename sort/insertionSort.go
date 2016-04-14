/**
 * codehackfox@gmail.com
 * 插入排序
 */

package main

import "fmt"

func main() {
    arr := [10]int {4,6,2,3,0,12,23,45,13,9}
    fmt.Println(arr)

    arr_len :=len(arr)
    for i:=1;i<arr_len;i++{
        temp:=arr[i]
        j:=i-1
        for ;j>=0 && arr[j]>temp;j--{
            arr[j+1]=arr[j]
        }
        arr[j+1]=temp
        fmt.Println(arr)
    }
}
