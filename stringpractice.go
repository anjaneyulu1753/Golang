package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "abcdefgh"
	name1 := "xswerabcdwd"
	new_slice := []string{}
	l := len(name)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if strings.Contains(name1, name[i:j+1]) {
				new_slice = append(new_slice, name[i:j+1])
			}
		}
	}
	fmt.Println(new_slice)
}
