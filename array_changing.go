package main

import "fmt"

func main() {
	weather := [3]string{"rainy", "sunny", "cloudy"}
	weather[1] = "stromy"
	fmt.Println(weather)
}
