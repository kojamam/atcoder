package main

import "fmt"

// 解説見た

func main() {
	var (
		n, cnt int
	)

	fmt.Scan(&n)

	for i := 11; i <= 55555; i += 5 {
		if isPrime(i) {
			fmt.Printf("%d ", i)
			cnt++
			if cnt == n {
				break
			}
		}
	}

}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
