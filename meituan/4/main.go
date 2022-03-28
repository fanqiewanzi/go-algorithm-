package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k1, k2 int
	fmt.Scan(&n, &k1, &k2)
	ball := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&ball[i])
	}
	var tmp int
	var p int
	sort.Ints(ball)
	for i := len(ball) - 1; i >= 0; i-- {
		tmp += ball[i]
		if tmp % k1 == 0 && tmp % k2 != 0 && tmp > p{
			p = tmp
		}
	}

	left, right := 0, len(ball) - 1
	tmp = 0
	var ans int
	for left < right {
		tmp += ball[left] + ball[right]
		if tmp < p {
			tmp -= ball[left]
			left++
		} else if tmp > p {
			tmp -= ball[right]
			right--
		} else {
			ans++
			tmp -= ball[left]
			left++
		}
	}
	fmt.Println(p, ans)
}

