package main

import "fmt"


func addBorder(picture []string) []string {
	final := []string{}
	asterics := ""
	for i:=0 ; i< len(picture[0]) + 2; i++ {
		asterics +="*"
	}
	final = append(final,asterics)
	for i:= 0; i < len(picture); i++ {
		str := "*" + picture[i] + "*"
		final = append(final, str)
	}
	final = append(final,asterics)
	return final
}

func main() {
	my_arr := []string{"abc", "ded"}
	fmt.Println(addBorder(my_arr))
}