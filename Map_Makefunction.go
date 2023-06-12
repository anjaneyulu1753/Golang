package main

import "fmt"

func main() {
	new_map := make(map[string]int)
	new_slice := make([]int, 5, 7)

	new_map["anji"] = 25
	new_map["dhanunjay"] = 26
	fmt.Println(new_map)
	fmt.Println(new_slice)

}
