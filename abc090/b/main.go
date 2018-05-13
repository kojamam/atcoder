package main

import "fmt"

func main() {
	var A, B, cnt int
	fmt.Scan(&A, &B)
	for i := A; i <= B; i++ {
		if i/10000 == i%10 && i/1000%10 == i/10%10 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
