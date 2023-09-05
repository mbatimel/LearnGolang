package main

import "fmt"

func main(){
	arguments := make(chan int)
	done := make(chan struct{})
	result := calculator(arguments, done)
	for i := 0; i < 10; i++ {
	   arguments <- 1
	}
	close(done)
	fmt.Println(<-result)
 }
 func calculator(arguments <-chan int, done <-chan struct{}) <-chan int{	
	resultChan := make(chan int)
	x:=0

	
	go func () {
		defer close(resultChan)
		for{
		select {
		case c := <- arguments:
			
			x += c
	
			
		case <-done:
			
			resultChan <- x
            return
		}
	}
	}()
	
	return resultChan

}