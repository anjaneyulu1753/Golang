package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "anjaneyulu"
	sub := "anji"
	sub1 := "ane"
	a := strings.Contains(name, sub)
	fmt.Println(a)
	b := strings.Contains(name, sub1)
	fmt.Println(b)
	c := strings.Replace(name, sub1, sub, 1)
	fmt.Println(c)
}
