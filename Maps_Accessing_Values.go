package main

import "fmt"

func main() {
	names_age := map[string]int{"Anji": 26, "Ram": 25, "Dhanunjay": 27, "Kishore": 28}
	for _, ages := range names_age {
		fmt.Println(ages)
	}
}
