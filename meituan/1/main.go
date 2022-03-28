package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	// 商品原价
	bPrices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&bPrices[i])
	}
	// 商品折扣价
	aPrices := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&aPrices[i])
	}
	// 满减规则的数量
	var m int
	fmt.Scan(&m)
	carr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&carr[i])
	}

	darr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&darr[i])
	}

	// 计算原价
	bPrices1 := make([]int, n)
	bPrices1[0] = bPrices[0]
	for i := 1; i < n; i++ {
		bPrices1[i] = bPrices1[i-1] + bPrices[i]
	}
	// 计算折扣价格
	aPrices1 := make([]int, n)
	aPrices1[0] = aPrices[0]
	for i := 1; i < n; i++ {
		aPrices1[i] += aPrices1[i-1] + aPrices[i]
	}
	for i := 0; i < n; i++ {
		min := bPrices1[i]
		// 计算满减后的最小值
		for j := m - 1; j >= 0; j-- {
			if bPrices1[i] >= carr[j] {
				min = bPrices1[i] - darr[j]
				break
			}
		}
		if min < aPrices1[i] {
			fmt.Printf("M")
		} else if aPrices1[i] < min {
			fmt.Printf("Z")
		} else {
			fmt.Printf("B")
		}
	}
}
