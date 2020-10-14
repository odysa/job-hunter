package model

type Company struct {
	Name        string
	Description string
}

type Job struct {
	Name        string
	Description string
	Requirement string
	Deadline    string
	Education   string
	SalaryLow   int
	SalaryHigh  int
	Location    string
	Company     Company
}
