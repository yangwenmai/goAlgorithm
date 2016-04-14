/**
 * codehackfox@gmail.com
 * 2016/4/14
 * 希尔排序
 */

 package main

 import "fmt"

 func main() {
     arr := [10]int {4,6,2,3,0,12,23,45,13,9}
     fmt.Println(arr)
     arr_len  := len(arr)

     for gap:=arr_len>>1;gap>0;gap>>=1{
         for i:=gap;i<arr_len;i++{
             temp :=arr[i]
             j :=i-gap
             for ;j>=0 && arr[j]>temp;j-=gap{
                 arr[j+gap] = arr[j]
             }
             arr[j+gap] = temp
         }
         fmt.Println(arr)
     }
 }
