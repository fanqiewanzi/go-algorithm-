package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

func combinationSum2(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int
	sort.Ints(candidates)
	var backTracking func(target, idx int)
	backTracking = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for i := idx; i < len(candidates) && target-candidates[i] >= 0; i++ {
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backTracking(target-candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}
	backTracking(target, 0)
	return ans
}
