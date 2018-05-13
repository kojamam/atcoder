package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	var n float64
	fmt.Scan(&a, &b, &c, &n)
	fmt.Println(a + b + c + max(a, b, c)*(int(math.Pow(2, n))-1))
}

func max(nums ...int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}
