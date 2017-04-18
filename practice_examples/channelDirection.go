package main 

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pongs chan<- string, pings <-chan string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	sendChannel := make(chan string)
	receiveChannel := make(chan string)
	go ping(sendChannel, "Ping")
	go pong(receiveChannel, sendChannel)
	fmt.Println(<-receiveChannel)
}