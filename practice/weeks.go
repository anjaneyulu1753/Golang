package main

import "fmt"

func main() {
    var dayOfWeek int = 4
    var dayName string

    switch dayOfWeek {
    case 1:
        dayName = "Monday"
        fallthrough
    case 2:
        dayName = "Tuesday"
        fallthrough
    case 3:
        dayName = "Wednesday"
        fallthrough
    case 4:
        dayName = "Thursday"
        fallthrough
    case 5:
        dayName = "Friday"
        fallthrough
    case 6:
        dayName = "Saturday"
        fallthrough
    case 7:
        dayName = "Sunday"
    }

    fmt.Printf("Day %d is %s\n", dayOfWeek, dayName)
}
