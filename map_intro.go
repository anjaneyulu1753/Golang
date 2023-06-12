package main

import "fmt"

func main() {
	new_map := map[string]int{"a": 1, "b": 2, "c": 2, "d": 4, "e": 5}
	new_map["c"] = 3
	fmt.Println(new_map)
	delete(new_map, "a")
	fmt.Println(new_map)
	for let, num := range new_map {
		fmt.Printf("%s:%d\n", let, num)
	}
	student := make(map[int]int)
	student[4] = 78
	student[7] = 98
	fmt.Println(student)
	// Access keys in a map
	for keys := range student {
		fmt.Println("keys:", keys)
	}
	// Access values in a map
	for _, values := range student {
		fmt.Println("values:", values)
	}
	// Access both keys & values in a map
	for keys, values := range student {
		fmt.Println(keys, ":", values)
	}
}
