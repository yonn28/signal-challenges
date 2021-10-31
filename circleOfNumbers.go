package main

import ("fmt"
"math"
)


func circleOfNumbers(n int, firstNumber int) int {
	calc := int(math.Floor(float64(n/2))) + firstNumber
	if calc > n {
		return calc - n
	}
	if calc == n {
		return 0
	}
	return calc 
}


func main() {
	fmt.Println(circleOfNumbers(10 ,7))
}