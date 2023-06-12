package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	var temp, rev, rem int
	temp = num
	rev = 0
	for num > 0 {
		rem = num % 10
		rev = rev*10 + rem
		num = num / 10
	}
	if temp == rev {
		fmt.Printf("%d is an palindrome.\n", temp)
	} else {
		fmt.Printf("%d is not an palindrome.\n", temp)
	}
}
