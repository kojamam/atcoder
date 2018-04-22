package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	spots := inputNumArray(N)

	spots = append([]int{0}, spots...)
	spots = append(spots, 0)

	var dirs []bool
	var diffs []int
	base := int(math.Abs(float64(spots[0])))
	for i := 0; i < N+1; i++ {
		if spots[i] < spots[i+1] {
			dirs = append(dirs, true)
		} else if spots[i] > spots[i+1] {
			dirs = append(dirs, false)
		} else {
			if i == 0 {
				dirs = append(dirs, true)
			} else {
				dirs = append(dirs, dirs[i-1])
			}
		}
		base += int(math.Abs(float64(spots[i] - spots[i+1])))

		if i > 0 {
			diff := 0
			if dirs[i] != dirs[i-1] {
				diff = int(math.Abs(float64(spots[i]-spots[i-1]))) + int(math.Abs(float64(spots[i]-spots[i+1]))) - int(math.Abs(float64(spots[i+1]-spots[i-1])))
			}
			diffs = append(diffs, diff)
		}
	}

	for i := 1; i < N+1; i++ {
		fmt.Println(base - diffs[i-1])
	}

}

func inputNumArray(length int) (nums []int) {
	for i := 0; i < length; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	return
}
