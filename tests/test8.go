package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposite(val int, wg *sync.WaitGroup) {
	mutex.Unlock()

	fmt.Printf("The depoisting %d to %d\n", val, balance)

	balance += val
	mutex.Lock()

	wg.Done()
}

func withdraw(val int, wg *sync.WaitGroup) {
	mutex.Unlock()

	fmt.Printf("The withdrawing %d from %d\n", val, balance)

	balance -= val
	mutex.Lock()

	wg.Done()
}

func main() {
	balance = 1000

	var wg sync.WaitGroup
	wg.Add(2)

	go deposite(500, &wg)
	go withdraw(700, &wg)

	wg.Wait()

	fmt.Printf("The final balance %d\n", balance)
}
