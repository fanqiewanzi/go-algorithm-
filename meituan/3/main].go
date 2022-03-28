package main

import "fmt"

func main() {
	var n, m1, m2 int
	fmt.Scan(&n, &m1, &m2)
	yxl := make([]int, m1)
	for i := 0; i < m1; i++{
		fmt.Scan(&yxl[i])
	}
	yxr := make([]int, m1)
	for i := 0; i < m1; i++{
		fmt.Scan(&yxr[i])
	}
	oldl := make([]int, m2)
	for i := 0; i < m2; i++{
		fmt.Scan(&oldl[i])
	}
	oldr := make([]int, m2)
	for i := 0; i < m2; i++{
		fmt.Scan(&oldr[i])
	}

	var ans int
	mp := make(map[int]struct{})
	for i := range yxl {
		for k := yxl[i]; k <= yxr[i]; k++ {
			mp[k] = struct{}{}
		}
	}

	for i := range oldl {
		for k := oldl[i]; k <= oldr[i]; k++ {
			if _, ok := mp[k]; ok {
				ans++
			}
		}
	}
	fmt.Println(ans)

}
