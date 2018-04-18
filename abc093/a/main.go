package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	line := readLine()
	flag := map[string]bool{"a": false, "b": false, "c": false}

	for i := 0; i < len(line); i++ {
		flag[line[i:i+1]] = true
	}

	if flag["a"] == true && flag["b"] == true && flag["c"] == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func readLine() string {
	sc.Scan()
	return sc.Text()
}
