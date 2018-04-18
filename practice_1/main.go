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

	numsStr := strings.Split(nextLine(), " ")
	numsStr = append(numsStr, strings.Split(nextLine(), " ")...)
	lastString := nextLine()

	sum := 0
	for i := 0; i < 3; i++ {
		num, _ := strconv.Atoi(numsStr[i])
		sum += num
	}

	fmt.Println(sum, lastString)

}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
