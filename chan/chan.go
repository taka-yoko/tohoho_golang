package main

import (
	"fmt"
	"time"
)

func funcA(chA chan<- string) {
	time.Sleep(3 * time.Second)
	chA <- "Finished"
}

func main() {
	chA := make(chan string)
	defer close(chA)
	go funcA(chA)
	msg := <-chA
	fmt.Println(msg)
}
