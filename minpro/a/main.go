package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if s[0:3] == "yah" && s[3] == s[4] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
