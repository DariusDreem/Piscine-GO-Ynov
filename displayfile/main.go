package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if len(os.Args[1:]) >= 2 {
		fmt.Println("Too many arguments")
	} else if len(os.Args[1:]) == 0 {
		fmt.Println("File name missing")
	}
	fmt.Printf("File contents: %s", content)
}
