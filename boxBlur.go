package main

import (
	"fmt"
	"math"
)

func makeSum(arr [][]int, row int, col int) int {
	var sum int = 0
	for i := row - 1; i < row+2; i++ {
		for j := col - 1; j < col+2; j++ {
			sum = sum + arr[i][j]
		}
	}
	return int(math.Floor(float64(sum / 9)))
}

func boxBlur(image [][]int) [][]int {
	result := make([][]int, len(image)-2)
	for i := 1; i < len(image)-1; i++ {
		result[i-1] = make([]int, len(image[i-1])-2)
		for j := 1; j < len(image[i])-1; j++ {
			result[i-1][j-1] = makeSum(image, i, j)
		}
	}
	return result
}

func main() {

	// fmt.Println(boxBlur([][]int{
	// 	{7, 4, 0, 1},
	// 	{5, 6, 2, 2},
	// 	{6, 10, 7, 8},
	// 	{1, 4, 2, 0}}))
	// fmt.Println(boxBlur([][]int{
	// 		{1, 1, 1},
	// 		{1, 7, 1},
	// 		{1, 1, 1}}))
	fmt.Println(boxBlur([][]int{
		{36, 0, 18, 9},
		{27, 54, 9, 0},
		{81, 63, 72, 45}}))
}
