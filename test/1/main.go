package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	ans := make([][]int, n*k)
	for i := range ans {
		ans[i] = make([]int, n*k)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for ni := i * k; ni < k*n; ni++ {
				for nj := j * k; nj < k*n; nj++ {
					ans[ni][nj] = matrix[i][j]
				}
			}
		}
	}

	for i := range ans {
		for j := range ans[i] {
			fmt.Printf("%d ", ans[i][j])
		}
		fmt.Println()
	}

}
