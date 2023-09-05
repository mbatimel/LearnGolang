package main
func main() {
	ch := make(chan int, 1)

go func(){

   //  work()
    ch <- 1
}()

<- ch


}