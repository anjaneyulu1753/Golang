
package main
import "fmt"

func main() {
	var num,remainder,temp int
	var reverse int = 0

	fmt.Print("Enter any positive integer : ")
	fmt.Scan(&num)

	temp=num

	// For Loop used in format of While Loop
	for num>0 {
		remainder = num%10
		reverse = reverse*10 + remainder
		num /= 10

	}

	if(temp==reverse){
		fmt.Printf("%d is a Palindrome",temp)
	}else{
		fmt.Printf("%d is not a Palindrome",temp)
	}

}