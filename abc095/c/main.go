package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	// input
	params := inputLineAsNumArray(5)
	pA := float64(params[0])
	pB := float64(params[1])
	pAB := float64(params[2])
	rA := params[3]
	rB := params[4]

	var a, b, ab int
	//created
	cA := 0
	cB := 0

	sum := 0.0
	if pAB < (pA+pB)/2 {
		min := min(rA, rB)
		sum += float64(min) * 2 * pAB
		cA += min
		cB += min

		ab += min * 2
	}

	if 2*pAB < pA {
		toCreate := (rA - cA)
		sum += float64(toCreate) * 2 * pAB
		cA += toCreate
		cB += toCreate

		ab += toCreate * 2
	}

	if 2*pAB < pB {
		toCreate := (rB - cB)
		sum += float64(toCreate) * 2 * pAB
		cA += toCreate
		cB += toCreate

		ab += toCreate * 2
	}

	if (rA - cA) > 0 {
		toCreate := (rA - cA)
		sum += float64(toCreate) * pA

		a += toCreate
	}

	if (rB - cB) > 0 {

		toCreate := (rB - cB)
		sum += float64(toCreate) * pB

		b += toCreate
	}

	fmt.Println(int(math.Ceil(sum)))
}

func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
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
