package main

import (
	"fmt"
	"strconv"
	"strings"
)

func evenDigitsOnly(n int) bool {
	number_string := strconv.Itoa(n)
	splited := strings.Split(number_string, "")
	isEven := true
	for _, value := range splited {
		elemnt, err := strconv.Atoi(value)
		if elemnt%2 != 0 && err == nil {
			isEven = false
		}
	}
	return isEven
}

func main() {
	fmt.Println(evenDigitsOnly(248622))
}
