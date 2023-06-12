package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "AnjAneYulu"
	name1 := "anjaneyulu"
	a := strings.ToLower(name)
	fmt.Println(a)
	b := strings.ToUpper(name)
	fmt.Println(b)
	c := strings.Count(name, "u")
	fmt.Println(c)
	d := strings.Split(name, "e")
	fmt.Println(d)
	e := strings.Index(name, "e")
	fmt.Println(e)
	f := strings.LastIndex(name, "e")
	fmt.Println(f)
	g := name == name1
	fmt.Println(g)
}
