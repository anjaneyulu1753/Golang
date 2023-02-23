package main

import (
    "fmt"
)

func main() {
    numbers := []int{2, 5, 8, 10, 15, 19, 20, 25, 30}

    var largestEven, largestOdd int

    for _, num := range numbers {
        if num%2 == 0 {
            if num > largestEven {
                largestEven = num
            }
        } else {
            if num > largestOdd {
                largestOdd = num
            }
        }
    }

    fmt.Println("Largest even number:", largestEven)
    fmt.Println("Largest odd number:", largestOdd)
}
