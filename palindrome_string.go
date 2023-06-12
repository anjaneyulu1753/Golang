package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	digit := 121
	name := strconv.Itoa(digit)
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
	name1 := ""
	l := len(name)
	for i := l - 1; i >= 0; i-- {
		name1 = name1 + name[i:i+1]
	}
	fmt.Println(name1)
	if name == name1 {
		fmt.Println("The given string: ", name, " is a palindrome")
	} else {
		fmt.Println("the given string is not a palindrome")
	}
}
