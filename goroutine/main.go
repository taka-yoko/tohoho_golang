package main

import (
	"fmt"
	"time"
)

func funcA() {
	for i := 0; i < 10; i++ {
		fmt.Println("A")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go funcA()
	for i := 0; i < 10; i++ {
		fmt.Println("M")
		time.Sleep(20 * time.Millisecond)
	}
}
