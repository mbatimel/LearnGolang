package main

import (
	"fmt"
	
)
func main() {                                   
	ch1, ch2 := make(chan int), make(chan int)  
	stop := make(chan struct{})                 
												
	r := calculator(ch1, ch2, stop)             
	//ch1 <- 3
	//ch2 <- 3   
	//close(stop)                                 
	fmt.Println(<-r)                            
												
 }   
func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int{
	resultChan := make(chan int)

	go func() {
		defer close(resultChan)

		select {
		case value := <-firstChan:
			resultChan <- value * value
		case value := <-secondChan:
			resultChan <- value * 3
		case <-stopChan:
			return
		}
	}()

	return resultChan
}