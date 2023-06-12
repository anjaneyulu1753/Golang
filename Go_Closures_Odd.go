package main

import "fmt"

func calculate() func() int {
	num := 1
	return func() int {
		num = num + 2
		return num
	}
}
func main() {
	odd := calculate()
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())

}
