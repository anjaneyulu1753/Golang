package main

import "fmt"

func main() {
    numbers := []int{2, 5, 8, 10, 15}

    fmt.Println("Original array:", numbers)

    if len(numbers) > 1 {
        numbers[0], numbers[len(numbers)-1] = numbers[len(numbers)-1], numbers[0]
    }

    fmt.Println("Array after swapping first and last elements:", numbers)
}
