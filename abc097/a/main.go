package main

import (
	"fmt"
	"math"
)

const debug = true

func main() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)

	if abs(A-C) <= D {
		fmt.Println("Yes")
		return
	}

	if A < B && B < C {
		if B-A <= D && C-B <= D {
			fmt.Println("Yes")
			return
		}
	}

	if A >= B && B >= C {
		if A-B <= D && B-C <= D {
			fmt.Println("Yes")
			return
		}
	}

	fmt.Println("No")
	return

}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}
