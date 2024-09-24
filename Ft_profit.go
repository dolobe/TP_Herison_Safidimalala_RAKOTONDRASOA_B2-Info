package main

import "fmt"

func Ft_profit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minimumPrice := prices[0]
	maximumProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minimumPrice {
			minimumPrice = prices[i]
		} else if prices[i]-minimumPrice > maximumProfit {
			maximumProfit = prices[i] - minimumPrice
		}
	}
	return maximumProfit
}

func main() {
	fmt.Println(Ft_profit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(Ft_profit([]int{7, 6, 4, 3, 1}))
}
