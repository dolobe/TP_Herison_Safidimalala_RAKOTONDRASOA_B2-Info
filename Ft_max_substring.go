package main

import "fmt"

func Ft_max_substring(s string) int {
	str := make(map[byte]int)
	maximum := 0
	first := 0
	for i := 0; i < len(s); i++ {
		if _, str1 := str[s[i]]; str1 {
			first = maxInt(str[s[i]], first)
		}

		maximum = maxInt(maximum, i-first+1)
		str[s[i]] = i + 1
	}
	return maximum
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(Ft_max_substring("abcabcbb"))
	fmt.Println(Ft_max_substring("bbbbb"))
}
