package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(time.Second * 1)
		wg.Done()
	}
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(25)
	go task("A", &waitGroup)
	go task("B", &waitGroup)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task anonymus is running\n", i)
			time.Sleep(time.Second * 1)
			waitGroup.Done()
		}
	}()
	waitGroup.Wait()
}
