package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Starting my program")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("Finishing goroutine")

		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Finishing my program")
}
