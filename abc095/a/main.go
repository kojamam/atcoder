package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	buf := make([]byte, 1000)
	sc.Buffer(buf, 1400000000)

	// input
	line := strings.Split(readLine(), "")

	price := 700
	for i := 0; i < 3; i++ {
		if line[i] == "o" {
			price += 100
		}
	}
	fmt.Println(price)
}

func readLine() string {

	sc.Scan()
	return sc.Text()
}
