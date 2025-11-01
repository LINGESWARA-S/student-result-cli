package main

import (
    "fmt"
    "os"
    "github.com/LINGESWARA-S/student-result-cli/models"
    "github.com/LINGESWARA-S/student-result-cli/result"
    "github.com/LINGESWARA-S/student-result-cli/utils"
)

func main() {
    if len(os.Args) > 1 {
        if os.Args[1] == "--example" {
            fmt.Println("Running example run for 2 students...")
            exampleRun()
            return
        }
    }

    fmt.Println("Student Result Processing CLI")
    fmt.Println("Press Ctrl+C to exit any time.")

    n := utils.AskInt("How many students to enter (1-100)? ", 1, 100)

    students := make([]models.Student, 0, n)
    for i := 0; i < n; i++ {
        fmt.Printf("\nEntering student %d:\n", i+1)
        s := utils.GetStudentFromInput()
        students = append(students, s)
    }

    results := result.ProcessBatch(students)

    fmt.Println("\n-- Results --")
    for _, r := range results {
        fmt.Printf("Roll: %s | Name: %s | Total: %d | Avg: %.2f | Grade: %s | Pass: %t\n",
            r.Roll, r.Name, r.Total, r.Average, r.Grade, r.Passed)
    }

    if utils.Confirm("Export results to file results.txt? (y/n): ") {
        err := utils.ExportResultsToFile("results.txt", results)
        if err != nil {
            fmt.Println("Failed to export:", err)
        } else {
            fmt.Println("Results exported to results.txt")
        }
    }
}

func exampleRun() {
    students := []models.Student{
        {Roll: "101", Name: "Alice", Marks: []int{85, 78, 92}},
        {Roll: "102", Name: "Bob", Marks: []int{42, 55, 38}},
    }

    results := result.ProcessBatch(students)
    fmt.Println("\n-- Results --")
    for _, r := range results {
        fmt.Printf("Roll: %s | Name: %s | Total: %d | Avg: %.2f | Grade: %s | Pass: %t\n",
            r.Roll, r.Name, r.Total, r.Average, r.Grade, r.Passed)
    }
}
