package main

import (
	"errors"
	"fmt"
)

func main() {
	funcA()
}

func funcA() (string, error) {
	var err error
	filename := ""
	data := ""

	filename, err = GetFileName()
	if err != nil {
		fmt.Println(err)
		goto Done
	}

	data, err = ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		goto Done
	}

Done:
	return data, err
}

func GetFileName() (string, error) {
	return "sample.txt", nil
}

func ReadFile(filename string) (string, error) {
	return "Hello world", errors.New("can't read file")
}
