package main

import "fmt"

func main() {
	languages := [...]string{"java", "python", "c", "go"}
	name := languages[0] + " " + languages[1]

	fmt.Println(name)
}
