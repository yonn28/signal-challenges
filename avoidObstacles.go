package main

import "fmt"

func isDivisible(inputArray []int, divisor int) bool {
	for i := 0; i < len(inputArray); i++ {
		if inputArray[i]%divisor == 0 {
			return false
		}
	}
	return true
}

func avoidObstacles(inputArray []int) int {
	isValidStep := false
	stepInit := 2
	for !isValidStep {
		isValidStep = isDivisible(inputArray, stepInit)
		if !isValidStep {
			stepInit++
		}
	}
	return stepInit
}

func main() {
	fmt.Println(avoidObstacles([]int{19, 32, 11, 23}))
	fmt.Println(avoidObstacles([]int{5, 3, 6, 7, 9}))
	fmt.Println(avoidObstacles([]int{2, 3}))
}
