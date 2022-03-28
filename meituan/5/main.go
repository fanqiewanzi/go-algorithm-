package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	do([]int{0})
	tmp := []int{0}
	for i := 0; i < n; i++ {
		tmp = do(tmp)
	}
	fmt.Println(tmp[k - 1])
}

func do(a []int) []int {
	var tmp []int
	for i := range a {
		tmp = append(tmp, a[i] + 1)
	}
	for i := range a {
		tmp = append(tmp, a[i])
	}
	for i := range a {
		tmp = append(tmp, a[i] + 2)
	}
	return tmp
}
