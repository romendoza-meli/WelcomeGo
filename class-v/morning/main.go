package main

import (
	"fmt"
	"time"
)

func main() {
	//create channel to communicate main and goroutines
	ch := make(chan int)
	go workAddition(ch, 1, 2)
	go workAddition(ch, 3, 4)
	go workAddition(ch, 5, 6)

	//time.Sleep(2 * time.Sencond)
	//fmt.Println("Waiting...")
	for {
		select {
		case result := <-ch:
			fmt.Println("Result", result)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			break
		}
	}
	result1 := <-ch

	fmt.Println("Waiting...")
	result2 := <-ch
	result3 := <-ch

	fmt.Println("Program finished", result1, result2, result3)
}

/* func workAddition(a, b int) {
	result := a + b
	time.Sleep(1 * time.Second)
	fmt.Println("Result:", result)
} */

func workAddition(ch chan int, a, b int) {
	result := a + b
	time.Sleep(1 * time.Second)

	// send result to channel
	ch <- result
}
