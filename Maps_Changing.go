package main

import "fmt"

func main() {

	names := map[string]int{"Anji": 26, "Ram": 25, "Dhanunjay": 27, "Kishore": 28}
	fmt.Println(names)
	names[`Anji`] = 25
	fmt.Println(names)
	names["Selva"] = 29
	names["Arun"] = 29
	fmt.Println(names)
	delete(names, "Ram")
	fmt.Println(names)
	fmt.Println(names)

}
