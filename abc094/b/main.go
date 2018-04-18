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

	line0 := strings.Split(nextLine(), " ")
	M, _ := strconv.Atoi(line0[1])
	X, _ := strconv.Atoi(line0[2])

	line1 := strings.Split(nextLine(), " ")
	costL := 0
	for i := 0; i < len(line1); i++ {
		num, _ := strconv.Atoi(line1[i])
		if num < X {
			costL += 1
		} else {
			break
		}
	}

	costR := M - costL
	if costL < costR {
		fmt.Println(costL)
	} else {
		fmt.Println(costR)
	}

}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
