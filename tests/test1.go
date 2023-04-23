package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func randSleep(wg *sync.WaitGroup, name string, limit int, sleep int) {
	defer wg.Done()
	for i := 1; i <= limit; i++ {
		fmt.Println(name, rand.Intn(i))
		time.Sleep(time.Duration(sleep * int(time.Second)))
	}
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1) // 2 goroutines

	go randSleep(wg, "first:", 10, 2)
	go randSleep(wg, "second:", 3, 2)

	wg.Wait()
}
