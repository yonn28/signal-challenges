package main

import (
	"fmt"
)

func isBlack(s string) bool {
	if (int64(s[0]) % 2 != 0 && int64(s[1]) % 2 != 0) || (int64(s[0]) % 2 == 0 && int64(s[1]) % 2 == 0){
		return true
	}
	return false
}

func chessBoardCellColor(cell1 string, cell2 string) bool {
	return isBlack(cell1) == isBlack(cell2)
 }


func main() {
	fmt.Println(chessBoardCellColor("A1","C3"))
}
