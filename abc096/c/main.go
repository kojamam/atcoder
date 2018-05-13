package main

import "fmt"

func main() {
	var (
		grid     [52][52]bool // t=>black, f=>white
		row, col int
	)

	fmt.Scan(&row, &col)

	for i := 1; i <= row; i++ {
		var line string
		fmt.Scan(&line)
		for j := 1; j <= col; j++ {
			if line[j-1] == '#' {
				grid[i][j] = true
			}
		}
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if grid[i][j] == true {
				if !(grid[i-1][j] || grid[i][j-1] || grid[i+1][j] || grid[i][j+1]) {
					fmt.Println("No")
					return
				}
			}
		}
	}

	fmt.Println("Yes")
}
