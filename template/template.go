package main

import "fmt"

const debug = true

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)
	nums := inputNums(B)

	if debug {
		fmt.Println(A, B, C)
		fmt.Println(nums)
	}
}

func inputNums(length int) (nums []int) {
	var num int
	for i := 0; i < length; i++ {
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}

func inputStrings(length int) (strings []string) {
	var str string
	for i := 0; i < length; i++ {
		fmt.Scanf("%s", &str)
		strings = append(strings, str)
	}
	return
}
