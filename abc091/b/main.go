package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	N := inputLineAsNumArray(1)[0]
	var blue []string
	for i := 0; i < N; i++ {
		blue = append(blue)
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

func inputLineAsStringArray(length int) []string {
	line := strings.Split(readLine(), " ")
	var str []string
	for i := 0; i < length; i++ {
		str = append(str, line[i])
	}
	return str
}

func readLine() string {
	sc.Scan()
	return sc.Text()
}
