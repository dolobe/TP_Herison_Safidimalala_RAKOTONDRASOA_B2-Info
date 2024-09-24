package main

import (
	"fmt"
	"sort"
)

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	value := 1
	sum := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= sum {
			value++
			sum = intervals[i][1]
		}
	}
	return len(intervals) - value
}

func main() {
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))
}
