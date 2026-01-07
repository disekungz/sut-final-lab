package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Employees struct {
	gorm.Model
	Name         string  `valid:"stringlength(2|80)"`
	Salary       float64 `valid:"range(15000|200000)~Salary must be between 15000 and 200000"`
	EmployeeCode string
}

func (e *Employees) Validate() (bool, error) {
	return govalidator.ValidateStruct(e)
}
