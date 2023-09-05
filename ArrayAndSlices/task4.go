package main

import (
	"fmt"
   
)

func main(){
 var n int 
 fmt.Scan(&n)
 arr := make([]int, n, n+1)
 for i := 0; i < n; i++{
    fmt.Scanf("%d", &arr[i])

 }
 for i := 0; i < n; i++{
    if i % 2 == 0 {
        fmt.Printf("%d ",arr[i])
    }

 }

}
