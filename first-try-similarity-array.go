package main

import (
	"fmt"
	"reflect"
)

func getCount( a []int ) map[int]int{
	count_a := make(map[int] int, 0)
	for _, value := range a {
		if _, ok := count_a[value]; ok {
			count_a[value] = count_a[value] + 1
		} else {
			count_a[value] = 1
		}
	}
	return count_a
}

func areSimilar(a []int, b []int) bool {
	count_a := getCount( a )
	count_b := getCount( b )
	res := reflect.DeepEqual(count_a, count_b)
	return res
}

func main() {
	fmt.Println(areSimilar([]int{1, 2, 2 },[]int{2, 1, 1}))
}