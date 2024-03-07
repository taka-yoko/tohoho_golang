package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

func funcA(a int, b ...int) {
	fmt.Printf("a=%d\n", a)
	for i, num := range b {
		fmt.Printf("b[%d]=%d\n", i, num)
	}
}

func main() {
	// fmt.Println(add(123, 456))
	// add, minus := addMinus(5, 3)
	// fmt.Println(add, minus)

	// _, minus2 := addMinus(7, 3)
	// fmt.Println(minus2)

	funcA(1, 2, 3, 9, 10)
}
