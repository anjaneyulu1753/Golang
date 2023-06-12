package main

import "fmt"

func main() {
	names_age := map[string]int{"anji": 25, "dhanunjay": 26, "kishore": 27, "selva": 28, "arun": 29}
	for name, age := range names_age {
		fmt.Printf("Name:%s and age: %d\n\n", name, age)
	}

}
