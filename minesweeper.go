package main

import "fmt"

func calculateTrues(matrix [][]bool, row int, col int) int {
	counter:= 0
	for i := row-1 ; i <= row+1; i++ {
		for j := col-1 ; j <= col+1; j++ {
			if ( i< len(matrix) && j <len(matrix[0]) ) && ( i>=0 && j>=0 ) && (i!=row || j!=col ) && matrix[i][j] {
				counter++
			}
		}
	}
	return counter
}

func minesweeper(matrix [][]bool) [][]int {
	result:= make([][]int, len(matrix))
    for i:=0 ; i< len(matrix) ; i++ {
        result[i] = make([]int, len(matrix[i]))
		for j:=0 ; j< len(matrix[i]) ; j++ {
			result[i][j] = calculateTrues(matrix, i, j)
		}
	}
	return result
}


func main() {
	fmt.Println(minesweeper([][]bool{
		{true,false,false}, 
        {false,true,false}, 
        {false,false,false},
	}))

	// fmt.Println(calculateTrues([][]bool{
	// 	{true,false,false}, 
    //     {false,true,false}, 
    //     {false,false,false},
	// },0,0))
}