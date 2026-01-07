package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func negative_employee_code(t *testing.T) {
	g := NewGomegaWithT(t)
	employees := Employees{
		Name:         "Testing test",
		Salary:       20000,
		EmployeeCode: "HR-120411",
	}
	ok, err := govalidator.ValidateStruct(employees)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"))
}
