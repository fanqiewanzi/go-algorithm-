package main

import (
	"fmt"
)

func main()  {
	var str string
	fmt.Scan(&str)
	ans := 1
	mp := make(map[string]struct{})
	var backTracking func(s string)
	backTracking = func(s string) {
		if len(s) < 2 {
			if _, ok := mp[string(s[0])]; !ok {
				mp[string(s[0])] = struct{}{}
				ans++
			}
			return
		}
		l, r := 0, 1
		for r < len(s) {
			if s[l] > s[r] {
				v := s[:l + 1] + s[r + 1:]
				if _,ok := mp[v]; ok {
					r++
					l++
					continue
				}
				ans++
				mp[v] = struct{}{}
				backTracking(v)
			} else if s[l] < s[r] {
				v := s[r:]
				if _,ok := mp[v]; ok {
					r++
					l++
					continue
				}
				mp[v] = struct{}{}
				ans++
				backTracking(v)
			} else {
				v1 := s[:l + 1] + s[r + 1:]
				v2 := s[r:]
				if _,ok := mp[v1]; ok {
					r++
					l++
					continue
				}
				mp[v1] = struct{}{}
				ans++
				backTracking(v1)
				if _,ok := mp[v2]; ok {
					r++
					l++
					continue
				}
				mp[v2] = struct{}{}
				ans++
				backTracking(v2)
			}
			r++
			l++
		}
	}
	backTracking(str)
	fmt.Println(ans % (1e9 + 7))
	fmt.Println(mp)
}
