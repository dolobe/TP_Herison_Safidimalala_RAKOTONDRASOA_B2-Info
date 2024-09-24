package main

import (
	"fmt"
)

func Ft_missing(nums []int) int {
	numbers := len(nums)
	somme := numbers * (numbers + 1) / 2
	some1 := 0
	for _, number := range nums {
		some1 += number
	}
	return somme - some1
}

func main() {
	fmt.Println(Ft_missing([]int{3, 1, 2}))
	fmt.Println(Ft_missing([]int{0, 1}))
	fmt.Println(Ft_missing([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
