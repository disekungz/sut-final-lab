package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func negative_sarary(t *testing.T) {
	g := NewGomegaWithT(t)
	employees := Employees{
		Name:         "Testing test",
		Salary:       1,
		EmployeeCode: "HR-1204",
	}
	ok, err := govalidator.ValidateStruct(employees)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
}
