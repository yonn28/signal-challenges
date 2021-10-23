package main

import (
	"fmt"
	"math"
)

func areEquallyStrong(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	yourLeftT := float64(yourLeft)
	yourRightT := float64(yourRight)
	friendsLeftT := float64(friendsLeft)
	friendsRightT := float64(friendsRight)
	return (math.Min(yourLeftT, yourRightT) == math.Min(friendsLeftT, friendsRightT) && math.Max(yourLeftT, yourRightT) == math.Max(friendsLeftT, friendsRightT))
}

func main() {
	fmt.Println(areEquallyStrong(10, 15, 15, 10))
}
