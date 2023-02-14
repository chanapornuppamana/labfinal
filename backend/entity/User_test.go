package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

type User struct {

	// Name     string `valid:"required~Name cannot be blank"`
	// Url      string `gorm:"uniqueIndex" valid:"url,required~url cannot be blank" `
	// Email    string `gorm:"uniqueIndex" valid:"email,required~Email not be blank"`
	// Username string `valid:"required~Username cannot be blank"`

	Name     string `valid:"alpha, required~Name not be blank"`
	Url      string `gorm:"uniqueIndex" valid:"url, required~url not be blank"`
	Email    string `gorm:"uniqueIndex" valid:"email, required~email not be blank"`
	Username string `valid:"alpha, required~Username not be blank"`
}

func TestCorrectAll(t *testing.T) {
	g := NewGomegaWithT(t)

	user := User{
		Name:     "chanaporn",
		Url:      "www.google.com",
		Email:    "chanaporn@gmail.com",
		Username: "nidnid",
	}
	ok, err := govalidator.ValidateStruct(user)

	g.Expect(ok).To(BeTrue())

	g.Expect(err).To(BeNil())

}

func TestNameNotBeBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	user := User{
		Name:     "", //false
		Url:      "www.google.com",
		Email:    "chanaporn@gmail.com",
		Username: "nidnid",
	}

	ok, err := govalidator.ValidateStruct(user)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Name not be blank"))
}
