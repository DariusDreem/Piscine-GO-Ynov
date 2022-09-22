package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	strInsert := ""
	resultlist := []string{}
	faitchier := false
	if len(arg) == 1 || arg[1] == "--help" || arg[1] == "-h" {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("   	 This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("   	 This flag will behave like a boolean, if it is called it will order the argument.")
		return
	}
	if len(arg) == 2 && arg[2][0] != '-' {
		fmt.Println(arg[1])
		return
	}
	if len(arg[1]) >= 4 && len(arg) >= 2 {
		a := arg[1][:3]
		if a == "-i=" {
			strInsert = arg[1][3:]
			arg = arg[2:]
			if arg[0][0] == '-' {
				arg[1] = arg[1] + strInsert
			} else {
				arg[0] = arg[0] + strInsert
			}
			faitchier = true
		}
	}
	if faitchier == false && (len(arg[1]) >= 10 && len(arg) >= 2) {
		a := arg[1][:9]
		if a == "--insert=" {
			strInsert = arg[1][9:]
			arg = arg[2:]

			if arg[0][0] == '-' {
				arg[1] = arg[1] + strInsert
			} else {
				arg[0] = arg[0] + strInsert
			}
		}
	}
	if faitchier == false {
		if arg[1] == "--order" || arg[1] == "-o" {
			for _, el := range arg[2] {
				resultlist = append(resultlist, string(el))
			}
			a := numberinorder(resultlist)
			arg[2] = ""
			for _, el := range a {
				arg[2] = arg[2] + el
			}
			fmt.Println(arg[2])
			return
		}
	}
	if arg[0] == "--order" || arg[0] == "-o" {
		for _, el := range arg[1] {
			resultlist = append(resultlist, string(el))
		}
		a := numberinorder(resultlist)
		arg[1] = ""
		for _, el := range a {
			arg[1] = arg[1] + el
		}
		fmt.Println(arg[1])
		return
	} else {
		fmt.Println(arg[0])
	}
}

func numberinorder(array []string) []string {
	for i := 0; i < len(array)-1; i++ {
		if array[i][0] > array[i+1][0] {
			array[i], array[i+1] = array[i+1], array[i]
			i = -1
		}
	}
	return array
}
