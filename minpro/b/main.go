package main

import (
	"fmt"
	"math"
)

func main() {
	var X, K int
	fmt.Scan(&X, &K)
	c := int(math.Pow10(K))
	fmt.Println(X/c*c + c)
}
