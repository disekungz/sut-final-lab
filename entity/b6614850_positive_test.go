package entity

import (
	"testing"

	"github.com/disekungz/sut-final-lab/entity"
	. "github.com/onsi/gomega"
)

func positive_test(t *testing.T) {
	t.Run("positve", func(t *testing.T) {
		g := NewGomegaWithT(t)
		employees := entity.Employees{
			Name:         "Testing test",
			Salary:       20000,
			EmployeeCode: "HR-1204",
		}
		ok, err := Validate(employees)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
