package main

type EventFor string

const (
	EventForTeacher = "teacher"
	EventForStudent = "student"
)

type RegisteredEvent struct {
	Email    string   `json:"email"`
	FullName string   `json:"full_name"`
	For      EventFor `json:"type"` // teacher or student
}
