package main

import "fmt"

func main() {
	x := 1
	y := 2

	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("y is greater than x")
	}

	mode := "running"

	switch mode {
	case "running":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("不明")
	}

	dayOfWeek := "sat"

	switch dayOfWeek {
	case "sat":
		fallthrough
	case "sun":
		fmt.Println("Holyday")
	default:
		fmt.Println("Weekday")
	}
}
