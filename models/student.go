package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name       string `json:"name" validate:"nonzero"`
	Student_id string `json:"student_id" validate:"len=8, regexp=^[0-9]*$"`
}

func FieldValidator(student *Student) error {
	if err := validator.Validate(student); err != nil {
		return err
	}

	return nil
}
