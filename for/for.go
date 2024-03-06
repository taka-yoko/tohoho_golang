package main

import "fmt"

func main() {
	x := 1
	y := 10

	for x < y {
		fmt.Println(x)
		x++
	}

	n := 0
	for {
		n++
		if n > 10 {
			break
		} else if n%2 == 1 {
			continue
		} else {
			fmt.Printf("%dは10以下の偶数\n", n)
		}
	}

	colors := [...]string{"red", "green", "blue"}
	for i, color := range colors {
		fmt.Printf("%d: %s\n", i, color)
	}
}
