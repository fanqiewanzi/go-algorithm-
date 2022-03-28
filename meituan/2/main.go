package main

import "fmt"

func main()  {
	var n, m int
	fmt.Scan(&n, &m)
	var str string
	fmt.Scan(&str)
	if m == 1 {
		var ans []byte
		for len(str) > 0 {
			var tmp int
			if len(str) % 2 != 0 {
				tmp = len(str) / 2 + 1
			} else {
				tmp = len(str) / 2
			}
			ans = append(ans, str[tmp - 1])
			str = str[:tmp - 1] + str[tmp:]
		}
		fmt.Println(string(ans))
	} else {
		var ans string
		for i := 0; i < n; i++ {
			if i % 2 != 0 {
				ans += string(str[i])
			} else {
				ans = string(str[i]) + ans
			}
		}
		fmt.Println(ans)
	}

}
