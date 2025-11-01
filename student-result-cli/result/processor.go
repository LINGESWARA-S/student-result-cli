package result

import "github.com/LINGESWARA-S/student-result-cli/models"

func ProcessBatch(students []models.Student) []models.Result {
    res := make([]models.Result, 0, len(students))
    for _, s := range students {
        r := EvaluateStudent(s)
        res = append(res, r)
    }
    return res
}

func EvaluateStudent(s models.Student) models.Result {
    total := 0
    for _, m := range s.Marks {
        total += m
    }
    avg := 0.0
    if len(s.Marks) > 0 {
        avg = float64(total) / float64(len(s.Marks))
    }

    grade := assignGrade(avg)
    passed := isPassed(s.Marks, avg)

    return models.Result{
        Roll:    s.Roll,
        Name:    s.Name,
        Total:   total,
        Average: avg,
        Grade:   grade,
        Passed:  passed,
    }
}

func assignGrade(avg float64) string {
    switch {
    case avg >= 90:
        return "A+"
    case avg >= 80:
        return "A"
    case avg >= 70:
        return "B"
    case avg >= 60:
        return "C"
    case avg >= 50:
        return "D"
    default:
        return "F"
    }
}

func isPassed(marks []int, avg float64) bool {
    if avg < 40 {
        return false
    }
    for _, m := range marks {
        if m < 35 {
            return false
        }
    }
    return true
}
