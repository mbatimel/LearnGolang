package main

import (
   "fmt"
   
//    "time"
)
// Я не дорешал
// у меня не получается я ебал того в рот

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	// Для контроля завершения всех горутин
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			x1, x2 := <-in1, <-in2

			// Каналы для результатов
			ch1, ch2 := make(chan int), make(chan int)

			// Запускаем вычисления в параллельных горутинах
			go func(x int, ch chan int) {
				ch <- fn(x)
			}(x1, ch1)

			go func(x int, ch chan int) {
				ch <- fn(x)
			}(x2, ch2)

			// Когда оба вычисления завершены, отправляем результат в out
			out <- <-ch1 + <-ch2
			done <- true
		}()
	}

	// Ждем завершения всех горутин
	for i := 0; i < n; i++ {
		<-done
	}
	close(out)
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

	for v := range out {
		fmt.Println(v)
	}

	close(in1)
	close(in2)
}
