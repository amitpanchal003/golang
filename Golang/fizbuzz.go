package main

import "fmt"

func fizzbuzz(cnt chan int, msg chan string) {
	for {
		i := <-cnt

		switch {
		case i%15 == 0:
			msg <- "FizzBuzz"
		case i%3 == 0:
			msg <- "fizz"
		case i%5 == 0:
			msg <- "Buzz"
		default:
			msg <- fmt.Sprintf("%d", i)
		}
	}

}

func main() {
	count := make(chan int)
	message := make(chan string)
	go fizzbuzz(count, message)
	for i := 1; i < 101; i++ {
		count <- i
		fmt.Println(<-message)
	}
}
