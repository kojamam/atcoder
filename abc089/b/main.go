package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var A string
		fmt.Scan(&A)
		if A == "Y" {
			fmt.Println("Four")
			return
		}
	}
	fmt.Println("Three")
}
