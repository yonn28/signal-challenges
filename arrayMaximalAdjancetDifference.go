package main

import "fmt"

func arrayMaximalAdjacentDifference(inputArray []int) int {
	max := 0
	for i := 1; i < len(inputArray); i++ {
		sub_calculated := inputArray[i-1] - inputArray[i]
		if sub_calculated < 0 {
			sub_calculated = -1 * sub_calculated
		}
		if sub_calculated > max {
			max = sub_calculated
		}
	}
	return max
}

func main() {
	fmt.Println(arrayMaximalAdjacentDifference([]int{10, 11, 13}))
}
