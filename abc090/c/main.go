package main

import "fmt"

func main() {
	var N, M, flipped int
	fmt.Scan(&N, &M)
	if N == 1 {
		N += 2
	}
	if M == 1 {
		M += 2
	}
	flipped = (N - 2) * (M - 2)

	fmt.Println(flipped)
}
