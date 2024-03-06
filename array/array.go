package main

import "fmt"

func main() {
	a1 := [3]string{}
	a1[0] = "Matsumoto"
	a1[1] = "Shinozuka"
	a1[2] = "Kuromati"

	fmt.Println(a1[0], a1[1], a1[2])

	a2 := [...]string{"Sakamoto", "Matsumoto", "Ogasawara"}

	fmt.Println(a2)

	// スライス
	a3 := []string{} // 要素の個数が不定
	a3 = append(a3, "Yashiki")
	a3 = append(a3, "Kato")
	a3 = append(a3, "Takagi")

	fmt.Println(a3)

	a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Println(len(a), cap(a))
	}

	bufa := make([]byte, 0, 1024)
}
