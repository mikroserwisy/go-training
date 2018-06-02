package main

func main() {
	channel := make(chan int)
	go func() {
		println(<- channel)
	}()

	channel <- 1
}
