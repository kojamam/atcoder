package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	buf := make([]byte, 1000)
	sc.Buffer(buf, 1400000000)

	// input
	nx := inputLineAsNumArray(2)
	N := nx[0]
	X := nx[1]

	sum := 0
	// var budgets []int
	minBudget := X
	for i := 0; i < N; i++ {
		budget := inputLineAsNumArray(1)[0]
		// budgets = append(budgets, budget)
		sum += budget
		if budget < minBudget {
			minBudget = budget
		}
	}

	fmt.Println(N + (X-sum)/minBudget)
}

func inputLineAsNumArray(length int) []int {

	line := strings.Split(readLine(), " ")
	var nums []int
	for i := 0; i < length; i++ {
		num, _ := strconv.Atoi(line[i])
		nums = append(nums, num)
	}
	return nums
}

func readLine() string {

	sc.Scan()
	return sc.Text()
}
