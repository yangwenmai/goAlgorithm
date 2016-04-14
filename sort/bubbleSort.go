/**
 * 冒泡排序
 * 1.初始化数组
 * 2.两两比较
 * codehackfox@gmail.com
 */

package main

import "fmt"

func main() {
    //声明并初始化数组
    arr := [10]int {4,6,2,3,0,12,23,45,13,9}
    fmt.Println(arr)

    l:=len(arr)
    for i:=0;i<l;i++ {
        flag:=true
        for k:=0;k<l-i-1;k++{
            if(arr[k]>arr[k+1]){
                flag = false
                arr[k],arr[k+1]=arr[k+1],arr[k]
            }
        }
        if flag {
            break
        }
        fmt.Println(arr)
    }
}
