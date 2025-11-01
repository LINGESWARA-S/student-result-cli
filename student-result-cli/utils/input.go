package utils

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
    "github.com/LINGESWARA-S/student-result-cli/models"
)

var reader = bufio.NewReader(os.Stdin)

func AskInt(prompt string, min, max int) int {
    for {
        fmt.Print(prompt)
        text, _ := reader.ReadString('\n')
        text = strings.TrimSpace(text)
        v, err := strconv.Atoi(text)
        if err == nil && v >= min && v <= max {
            return v
        }
        fmt.Printf("Please enter a valid integer between %d and %d.\n", min, max)
    }
}

func Confirm(prompt string) bool {
    for {
        fmt.Print(prompt)
        text, _ := reader.ReadString('\n')
        t := strings.ToLower(strings.TrimSpace(text))
        if t == "y" || t == "yes" {
            return true
        }
        if t == "n" || t == "no" {
            return false
        }
        fmt.Println("Please type 'y' or 'n'.")
    }
}

func GetStudentFromInput() models.Student {
    fmt.Print("Enter Roll: ")
    roll, _ := reader.ReadString('\n')
    roll = strings.TrimSpace(roll)

    fmt.Print("Enter Name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    sub := AskInt("Number of subjects (1-10): ", 1, 10)
    marks := make([]int, sub)
    for i := 0; i < sub; i++ {
        prompt := fmt.Sprintf("Enter marks for subject %d (0-100): ", i+1)
        marks[i] = AskInt(prompt, 0, 100)
    }

    return models.Student{
        Roll:  roll,
        Name:  name,
        Marks: marks,
    }
}

func ExportResultsToFile(filename string, results []models.Result) error {
    f, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer f.Close()

    for _, r := range results {
        line := fmt.Sprintf("Roll: %s | Name: %s | Total: %d | Avg: %.2f | Grade: %s | Pass: %t\n",
            r.Roll, r.Name, r.Total, r.Average, r.Grade, r.Passed)
        _, err := f.WriteString(line)
        if err != nil {
            return err
        }
    }
    return nil
}
