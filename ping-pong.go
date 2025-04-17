package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go func() {
		i := 1
		for i < 10 {
			time.Sleep(time.Duration(i) * time.Second)

			ping <- "ping"
			i = i + 2
		}
	}()
	go func() {
		i := 2
		for i <= 10 {
			time.Sleep(time.Duration(i) * time.Second)

			pong <- "pong"
			i = i + 2
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-ping:
			fmt.Println(msg1)
		case msg2 := <-pong:
			fmt.Println(msg2)
		}
	}
}
