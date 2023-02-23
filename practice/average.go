package main

import "fmt"

func main() {
    numbers := []int{2, 5, 8, 10, 15}

    var sum, avg float64

    for _, num := range numbers {
        sum += float64(num)
    }

    if len(numbers) > 0 {
        avg = sum / float64(len(numbers))
    }

    fmt.Printf("Average of the array: %.2f", avg)
}
