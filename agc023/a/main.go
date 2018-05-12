package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var n int
	var sameSums = make(map[int]int) //要素0からの部分和が同じものの個数を数え上げたもの

	// 要素0の初期化
	fmt.Scan(&n)
	prevk := n
	sameSums[n]++

	//空な部分列の和は0
	sameSums[0]++

	for i := 1; i < N; i++ {
		fmt.Scan(&n)
		k := prevk + n
		sameSums[k]++
		prevk = k
	}

	cnt := 0
	for _, v := range sameSums {
		cnt += combination(v, 2)
	}

	fmt.Println(cnt)
}

func combination(n, r int) int {
	a := 1
	b := 1
	for i := 0; i < r; i++ {
		a *= n - i
		b *= r - i
	}
	return a / b
}
