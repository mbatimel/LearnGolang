package main

import (
	"fmt"
	"time"
)
func main() {
	var m =map[int]int{}
	var arr [10]int
	
	for i:=0;i<10;i++{
	  fmt.Scan(&arr[i])
	  m[arr[i]] = 0
	}
  
	for i:=0;i<10;i++{
	  if m[arr[i]] == 0{
		m[arr[i]] = work(arr[i])
	  }
	  fmt.Print(m[arr[i]]," ")
	}
}
func work(n int) int {
	if n > 3 {
	   time.Sleep(time.Millisecond * 500)
	   return n + 1
	} else {
	   time.Sleep(time.Millisecond * 500)
	   return n - 1
	}
 }