package main

import "fmt"

type Person struct {
	name string
}

func (p Person) ToString() string {
	return p.name
}
func (p Person) PrintOut() {
	fmt.Println(p.ToString())
}

type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}
func (b Book) PrintOut() {
	fmt.Println(b.ToString())
}

func main() {
	a1 := Person{name: "司馬遼太郎"}
	a2 := Book{title: "新記太閤記"}
	a1.PrintOut()
	a2.PrintOut()
}
