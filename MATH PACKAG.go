package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("square root of 25 is: ", math.Sqrt(25))
	fmt.Println("cube root of 64 is: ", math.Cbrt(64))
	fmt.Println("larger number:", math.Max(34, 43))
	fmt.Println("smaller number: ", math.Min(37, 494))
	fmt.Println("remainder is: ", math.Mod(23, 3))

}
