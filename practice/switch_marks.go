package main

import "fmt"

func main() {
    var grade string
    marks := 85

    switch marks {
        case 90:
            grade = "A"
        case 80:
            grade = "B"
        case 70:
            grade = "C"
        case 60:
            grade = "D"
        default:
            grade = "F"
    }

    fmt.Printf("Marks: %d\nGrade: %s", marks, grade)
}
