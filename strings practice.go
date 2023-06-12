package main

import "fmt"

func main() {
	name := "anjaneyulu"
	new_slice := []string{}
	for j := 1; j < len(name)+1; j++ {
		for i := 0; i < len(name)-(j-1); i++ {
			new_slice = append(new_slice, name[i:i+j])
		}
	}

	fmt.Println(new_slice)
}
