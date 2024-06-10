package student

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name"`
	Student_id   string `json:"student_id"`
}

var Students []Student
