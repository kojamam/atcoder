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

	// resize buffer
	buf := make([]byte, 1000)
	sc.Buffer(buf, 200000000000) //頭悪い

	// input
	N := inputLineAsNumArray(1)[0]
	nums, max := inputLineAsNumArrayWithMax(N)

	// combination
	med := int(max / 2)

	diff_idx := 0
	diff := max
	for i := 0; i < N; i++ {
		if nums[i] == max {
			continue
		}
		if int(math.Abs(float64(med-nums[i]))) < diff {
			diff_idx = i
			diff = int(math.Abs(float64(med - nums[i])))
		}
	}
	fmt.Println(max, nums[diff_idx])
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

func inputLineAsNumArrayWithMax(length int) ([]int, int) {

	line := strings.Split(readLine(), " ")
	var nums []int
	max := 0
	for i := 0; i < length; i++ {
		num, _ := strconv.Atoi(line[i])
		nums = append(nums, num)
		if max < num {
			max = num
		}
	}
	return nums, max
}

func readLine() string {

	sc.Scan()
	return sc.Text()
}
