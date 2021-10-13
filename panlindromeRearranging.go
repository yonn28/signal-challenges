package main

import (
	"fmt"
)

func getCount( a string ) map[rune]int{
	count_a := make(map[rune] int, 0)
	for _, value := range a {
		if _, ok := count_a[value]; ok {
			count_a[value] = count_a[value] + 1
		} else {
			count_a[value] = 1
		}
	}
	return count_a
}

func palindromeRearranging(inputString string) bool {
	secondEven := true
	countEven := 0
	for _, v := range getCount(inputString) {
		if ( v % 2 == 1 && secondEven ) {
			if(countEven == 1) {
				secondEven = false
			}
			countEven++
		}
	}
	return secondEven
}



func main() {
	fmt.Println(palindromeRearranging("aabb") ) /*true*/
	fmt.Println(palindromeRearranging("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabc") )/*false*/
	fmt.Println(palindromeRearranging("abbcabb") )/*true*/
	fmt.Println(palindromeRearranging("zyyzzzzz") )/*true*/
	fmt.Println(palindromeRearranging("z") )/*true*/
	fmt.Println(palindromeRearranging("abdhuierf") )/*false*/
	fmt.Println(palindromeRearranging("abdhuierf") )/*false*/
}