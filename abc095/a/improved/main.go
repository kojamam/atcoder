package main

import "fmt"

func main() {
	var toppingList string
	fmt.Scan(&toppingList)
	sum := 700
	for i := 0; i < 3; i++ {
		if toppingList[i] == 'o' {
			sum += 100
		}
	}
	fmt.Println(sum)
}
