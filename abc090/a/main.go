package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		var str string
		fmt.Scan(&str)
		fmt.Printf("%c", str[i])
	}
	fmt.Print("\n")
}
