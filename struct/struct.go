package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}

type Person2 struct {
	height float32
	weight float32
}

func NewPerson(height, weight float32) *Person2 {
	return &Person2{height: height, weight: weight}
}

func main() {
	var p1 Person
	p1.SetPerson("Zakoshi", 46)
	name, age := p1.GetPerson()
	fmt.Printf("%s(%d)\n", name, age)

	// 初期化
	a1 := Person{"Yamada", 26}
	a2 := Person{name: "Tanaka", age: 29}
	fmt.Println(a1.name)
	fmt.Println(a2.name)

	p := NewPerson(182.3, 120.5)
	fmt.Println(p.height)
}
