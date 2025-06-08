package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d recebeu %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	data := make(chan int)
	qntWorkwers := 100000

	for i := 0; i <= qntWorkwers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100000000; i++ {
		data <- i
	}
}
