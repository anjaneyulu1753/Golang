package main

import "fmt"

func main() {
	var n, sum int = 100, 0
	for i := 0; i <= n; i++ {
		fmt.Print(i, ",")
		sum += i
	}
	fmt.Println("The sum is: ", sum)
}
