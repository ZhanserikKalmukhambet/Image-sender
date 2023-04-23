package main

import (
	"fmt"
	"sync"
)

func writeChannel(wg *sync.WaitGroup, limitchannel chan int, stop int) {
	defer wg.Done()
	for i := 1; i <= stop; i++ {
		limitchannel <- i
	}
}

func readChannel(wg *sync.WaitGroup, limitchannel chan int, stop int) {
	defer wg.Done()
	for i := 1; i <= stop; i++ {
		fmt.Println(<-limitchannel)
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	limitchannel := make(chan int)
	defer close(limitchannel)

	go writeChannel(wg, limitchannel, 3)
	go readChannel(wg, limitchannel, 3)

	wg.Wait()
}