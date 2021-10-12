package main

import (
	"fmt"
)

func arrayChange(inputArray []int) int {
	count := 0
	for i:= 1 ; i < len(inputArray) ; i++ {
		if( inputArray[i-1] >= inputArray[i] ){
			for inputArray[i-1] >= inputArray[i] {
				inputArray[i] = inputArray[i] + 1
				count++
			}
		}
	}
	return count
}


func main() {
	fmt.Println(arrayChange([]int{-1000, 0, -2, 0}))
}