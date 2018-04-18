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
	// input
	nums := inputLineAsNumArray(3)
	minIdx, midIdx, maxIdx := getOrder3(nums)

	// count
	count := 0

	// first step
	numFirstStep := nums[maxIdx] - nums[midIdx]

	for i := 0; i < numFirstStep; i++ {
		nums[minIdx]++
		nums[midIdx]++
		count++
	}

	// second step
	for nums[minIdx] < nums[maxIdx] {
		nums[minIdx] += 2
		count++
	}

	// third step (if required)
	if nums[minIdx] > nums[maxIdx] {
		nums[midIdx]++
		nums[maxIdx]++
		count++
	}

	fmt.Println(count)
}

func getOrder3(nums []int) (int, int, int) {
	max := nums[0]
	min := nums[0]
	var minIdx, midIdx, maxIdx int

	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
		if nums[i] < min {
			min = nums[i]
			minIdx = i
		}
	}
	switch minIdx + maxIdx {
	case 1:
		midIdx = 2
	case 2:
		midIdx = 1
	case 3:
		midIdx = 0
	default:
		midIdx = 0
	}

	return minIdx, midIdx, maxIdx
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
