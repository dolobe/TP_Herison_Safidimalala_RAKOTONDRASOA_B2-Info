package main

import "fmt"

func Ft_profit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))
}
