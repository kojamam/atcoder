package main

import (
	"fmt"
	"math"
)

/* main */

func main() {
	var N int
	fmt.Scan(&N)

	fmt.Println(N)
}

/* IO */

func scanNums(length int) (nums []int) {
	var num int
	for i := 0; i < length; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}

func scanStrings(length int) (strings []string) {
	var str string
	for i := 0; i < length; i++ {
		fmt.Scanf("%s", &str)
		strings = append(strings, str)
	}
	return
}

/* intのmath */

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func pow(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}

func min(nums ...int) int {
	res := math.MaxFloat64
	for i := 0; i < len(nums); i++ {
		res = math.Min(res, float64(nums[i]))
	}
	return int(res)
}

func max(nums ...int) int {
	var res float64
	for i := 0; i < len(nums); i++ {
		res = math.Max(res, float64(nums[i]))
	}
	return int(res)
}

/* 文字列 */

func strSearch(a []string, b string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == b {
			return true
		}
	}
	return false
}
