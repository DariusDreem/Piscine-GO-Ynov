package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	file, err := os.ReadFile("quest8.txt")

	if len(arg) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(arg) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Print(string(file))
}
