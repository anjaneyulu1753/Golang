package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Anjaneyulu Gudapati"
	name1 := `ram`
	sub := "ane"
	fmt.Println(name)
	fmt.Println(name1)
	fmt.Printf("%c\n", name[0])
	fmt.Println(name[1:3])
	fmt.Printf("the length of the given string %s is:%d\n ", name, len(name))
	a := name1 + " " + name
	fmt.Println(a)

	fmt.Println(strings.Compare(name, name1))
	fmt.Println("the given substring present in main string is: ", strings.Contains(name, sub))
	fmt.Println(strings.Replace(name, "u", "m", 4))

	fmt.Println("string in lower case: ", strings.ToLower(name))
	fmt.Println("string in upper case: ", strings.ToUpper(name))
	fmt.Println("after split: ", strings.Split(name, ""))
	b := strings.Count(name, "n")
	fmt.Println(b)
	c := strings.Index(name, "a")
	fmt.Println(c)
}
