package main

import "fmt"

func main() {
	var channel1 chan int
	fmt.Println(channel1)

	channel2 := make(chan int)
	fmt.Println(channel2)
}
