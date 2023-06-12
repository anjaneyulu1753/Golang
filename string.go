package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string = "anjaneyulu"
	var name1 string = "anjaneyulu"
	fmt.Println(strings.Compare(name, name1))

	fmt.Println(name[4:])
}
