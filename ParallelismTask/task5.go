package main

import (
	"fmt"
	
)


func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	for i := 0; i < n; i++ {
		go func() {
			
			x1, x2 := <-in1, <-in2

			
			
			go func() {
				x1 = fn(x1)
			}()

			go func() {
				x2 = fn(x2)
			}()

			
			out <- x1+x2
			
		}()
	}
	
}

func main() {
	fn := func(x int) int {
		return x * x
	}

	in1 := make(chan int, 5)
	in2 := make(chan int, 5)
	out := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		in1 <- i
		in2 <- i + 5
	}

	merge2Channels(fn, in1, in2, out, 5)
	
	for i:=1; i <= 5; i++ {
		fmt.Println(<-out)
	}
	

	
}
