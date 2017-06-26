package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	var m int

	io := bufio.NewReader(os.Stdin)

	fmt.Fscan(io, &n)
	fmt.Fscan(io, &m)

	arr := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(io, &arr[i])
	}

	fmt.Println(count(arr, m, n))

}

func count(S []int, m int, n int) int {
	var x, y int

	table := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		table[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		table[0][i] = 1
	}

	for i := 1; i < n+1; i++ {
		for j := 0; j < m; j++ {

			if i-S[j] >= 0 {
				x = table[i-S[j]][j]
			} else {
				x = 0
			}

			if j >= 1 {
				y = table[i][j-1]
			} else {
				y = 0
			}

			table[i][j] = x + y
		}
	}
	return table[n][m-1]
}
