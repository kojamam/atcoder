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

	nums := inputLineAsNumArray(3)
	A := nums[0]
	B := nums[1]
	C := nums[2]
	if A+B >= C {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
