package main

import "fmt"

func main() {
	var workArray [10]uint8
	var indexArray [6]int
	 for i := 0; i < 10; i++ {
		 fmt.Scan(&workArray[i])
	 }
	 for i := 0; i < 6; i++ {
		 fmt.Scan(&indexArray[i])
	 }
 
	 for i := 0; i <= 4; i+=2 {
		 j := workArray[indexArray[i]]
		 workArray[indexArray[i]] = workArray[indexArray[i+1]]
		 workArray[indexArray[i+1]] = j
	 }
	 for i := 0; i < 10; i++ {
		 
			fmt.Printf("%d ",workArray[i]) 
		 
	 }
}