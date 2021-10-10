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

func swap( input []int, i int, j int) []int {
	input[i], input[j] = input[j], input[i]
	return  input
}

func areSimilar(a []int, b []int) bool {
    count_a := getCount( a )
	count_b := getCount( b )
	res := reflect.DeepEqual(count_a, count_b)
	temp := b
	if( res ){
		for i:= 0 ; i < len(a); i++ {
			for j:= 0 ; j < len(a); j++ {

				temp = swap(temp, i, j)
				isEqual := reflect.DeepEqual(temp, a)

				if( isEqual ) {
					return isEqual;
					break
				}
				temp = swap(temp, i, j)
			}
		}
	}
	return false
}

func main() {
	fmt.Println(areSimilar([]int{2, 3, 1},[]int{1, 3, 2}))
}