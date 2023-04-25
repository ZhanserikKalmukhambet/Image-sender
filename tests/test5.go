package main

import "fmt"

func main() {
	letter := make(chan string, 2) // 2 - length of buffered channel

	letter <- "hello"
	fmt.Println(<-letter)

	letter <- "world"
	letter <- "!"

	fmt.Println(<-letter)
	fmt.Println(<-letter)

	// buffered channels will not be block in the process, while buffered do it

}
