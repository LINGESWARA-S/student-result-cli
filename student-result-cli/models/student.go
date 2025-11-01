package models

type Student struct {
    Roll  string
    Name  string
    Marks []int
}

type Result struct {
    Roll    string
    Name    string
    Total   int
    Average float64
    Grade   string
    Passed  bool
}
