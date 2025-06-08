package main

import "fmt"

// Thread 1
func main() {
	channel := make(chan string)

	//Thread 2
	go func() {
		channel <- "Hello World!" //Canal fica cheio
	}()

	//Thread 1
	msg := <-channel //Canal esvazia
	fmt.Println(msg)
}
