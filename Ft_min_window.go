package main

import "fmt"

func Ft_min_window(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	tab := make(map[byte]int)
	for i := range t {
		tab[t[i]]++
	}

	first := 0
	fin := 0
	min := len(s) + 1
	count := len(t)
	start := 0

	for fin < len(s) {
		if tab[s[fin]] > 0 {
			count--
		}
		tab[s[fin]]--
		fin++

		for count == 0 {
			if fin-first < min {
				min = fin - first
				start = first
			}
			tab[s[first]]++
			if tab[s[first]] > 0 {
				count++
			}
			first++
		}
	}
	if min == len(s)+1 {
		return ""
	}
	return s[start : start+min]

}

func main() {
	fmt.Println(Ft_min_window("ADOBECODEBANC", "ABC"))
	fmt.Println(Ft_min_window("a", "aa"))
}
