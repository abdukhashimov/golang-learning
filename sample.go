package main

import "fmt"

func main() {
	messages := make(chan string, 2)
	go func() {
		messages <- "ping"
		messages <- "pinasdg"
		messages <- "pingasd"
		messages <- "pinasdg"
		messages <- "asdping"
		messages <- "ping"
	}()
	msg := <-messages
	msg2 := <-messages
	msg3 := <-messages
	fmt.Println(msg)
	fmt.Println(msg2)
	fmt.Println(msg3)
	fmt.Println("Hello World")
}
