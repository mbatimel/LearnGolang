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
 fmt.Print(arr[3])

}
