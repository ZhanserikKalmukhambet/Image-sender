package main

import (
	"fmt"
	"time"
)

func main() {
	mess1 := make(chan string)
	mess2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			mess1 <- "Прошло пол секунды"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			mess2 <- "Прошло 2 секунды"
		}
	}()

	for {
		select {
		case msg := <-mess1:
			fmt.Println(msg)
		case msg := <-mess2:
			fmt.Println(msg)
		}
	}
}
