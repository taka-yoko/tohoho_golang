package main

import "fmt"

type Printable interface {
	ToString() string
}

func PrintOut(a interface{}) {
	q, ok := a.(Printable)
	if ok {
		fmt.Println(q.ToString())
	} else {
		fmt.Println("Not printable.")
	}
}

type Person struct {
	name string
}

func (p Person) ToString() string {
	return p.name
}

type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}

func funcA(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Printf("%t\n", a)
	case int:
		fmt.Printf("%d\n", a)
	case string:
		fmt.Printf("%s\n", a)
	}
}

type any interface{}
type dict map[string]any

func main() {
	a1 := Person{name: "司馬遼太郎"}
	a2 := Book{title: "新記太閤記"}
	PrintOut(a1)
	PrintOut(a2)

	a3 := true
	a4 := 456
	a5 := "hoge"
	funcA(a3)
	funcA(a4)
	funcA(a5)

	p1 := dict{
		"name": "Tsutsugo",
		"age":  31,
		"address": dict{
			"zip": "123-4567",
			"tel": "012-3456-7890",
		},
	}
	name := p1["name"]
	zip := p1["address"].(dict)["zip"]
	fmt.Println(name, zip)
}
