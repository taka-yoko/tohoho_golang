package main

import "fmt"

func main() {
	var a1 int  // int型変数a1
	var p1 *int // int型へのポインタ変数p1

	p1 = &a1  // p1にa1のポインタを設定
	*p1 = 123 // ポインタの中身(a1)に123を代入
	fmt.Println(a1)

	var a2 int = 123
	var a3 int = 123
	fn(a2, &a3)
	fmt.Println(a2, a3) // 123 456

	a4 := Person{"tanaka", 26}
	p2 := &a4
	fmt.Println(a4.name)
	fmt.Println((*p2).name)
	fmt.Println(p2.name)
}

func fn(b1 int, b2 *int) {
	b1 = 456
	*b2 = 456
}

type Person struct {
	name string
	age  int
}
