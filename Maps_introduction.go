package main

import "fmt"

func main() {
	new_map := map[string]float32{"golang": 89.9, "java": 90, "python": 91}
	fmt.Println(new_map)
	fmt.Println(new_map["golang"])
	fmt.Println(new_map["python"])
	fmt.Println(new_map["java"])
}
