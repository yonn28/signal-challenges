package main

import (
	"fmt"
	"regexp"
)

func variableName(name string) bool {
	validator := regexp.MustCompile("^[A-Za-z_]{1}[a-z_A-Z0-9]*$")
	return validator.MatchString(name)
}

func main() {
	fmt.Println(variableName("var_1__Int"))
	fmt.Println(variableName("qq-q"))
	fmt.Println(variableName("2w2"))
	fmt.Println(variableName("va[riable0"))
}
