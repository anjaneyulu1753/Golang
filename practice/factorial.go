package main
import "fmt"

func main() {
	var num int
	factorial:=1

	fmt.Print("Enter any positive integer : ")
	fmt.Scan(&num)
	for i := 1; i < num+1; i++ {
	factorial=factorial*i
	}
	fmt.Println("the factorial sum is:", factorial)

}