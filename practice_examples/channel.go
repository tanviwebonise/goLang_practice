package main

import (
    "fmt"
    "time"
)

func makeCakeAndSend(cs chan string) {
	cakeName := "Strawberry Cake"
	for i := 1; i <= 3; i++ {
	    fmt.Println("Making a cake and sending ...", cakeName, i)
	    cs <- cakeName //send a strawberry cake
	}
}

func receiveCakeAndPack(cs chan string) {
	for i := 1; i <= 3; i++ {
	    s := <-cs //get cake is on the channel
	    fmt.Println("Packing received cake: ", s, i)
	}
}

func main() {
    cs := make(chan string, 3)//buffred channel
    go makeCakeAndSend(cs)
    go receiveCakeAndPack(cs)
    time.Sleep(time.Second * 1)
}