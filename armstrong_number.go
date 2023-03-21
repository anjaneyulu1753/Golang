// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello,")
	var num, rem, temp int
	var sum = 0
	fmt.Print("Enter the three digit number: ")
	fmt.Scanln(&num)
	temp = num
	for {
		rem = temp % 10
		sum += rem * rem * rem
		temp = temp / 10
		if temp == 0 {
			break
		}
	}
	if sum == num {
		fmt.Printf("%d is an armstrong number.", num)
	} else {
		fmt.Printf("%d is not an armstrong number.", num)
	}
}
