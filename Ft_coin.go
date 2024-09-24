package ymmersion

import (
	"fmt"
)

func minimum(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func Ft_coin(coins []int, amount int) int {
	tableau := make([]int, amount+1)
	for i := range tableau {
		tableau[i] = amount + 1
	}
	tableau[0] = 0

	for i := 1; i <= amount; i++ {
		for _, piece := range coins {
			if piece <= i {
				tableau[i] = minimum(tableau[i], tableau[i-piece]+1)
			}
		}
	}
	if tableau[amount] < amount {
		fmt.Println(-1)
		return -1
	}

	fmt.Println(tableau[amount])
	return tableau[amount]
}
