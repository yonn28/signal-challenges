package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isIPv4Address(inputString string) bool {

	result := strings.Split(inputString, ".")
	for _, v := range result {
		start_with_zero := strings.HasPrefix(v, "0")
		transform, err := strconv.Atoi(v)
		if transform > 255 || len(result) != 4 || err != nil || (start_with_zero && v != "0") {

			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isIPv4Address("0.254.255.0"))
}
