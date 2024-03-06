package main

func main() {
	var a1 int = 123
	var a2 = 123
	a3 := 123

	var (
		a1 int = 123
		a2 int = 456
	)

	var name, age = "Yamada", 26

	const foo = 100
	const (
		foo2 = 100
		bar  = 200
	)

	// 型に別名をつける
	type UtcTime string
	type JstTime string

	var t1 UtcTime = "00:00:00"
	var t2 JstTime = "09:00:00"
	// t1 = t2 // 型が異なるので代入エラーになる

	// 型変換
	var a5 uint16 = 1234
	var a6 uint32 = uint32(a5)
}
