package models

import (
	"final-project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserName  string     `gorm:"not null" json:"full_name" form:"full_name" `
	Email     string     `gorm:"not null" json:"email" form:"email" valid:"required-Your email is required,email-Invalid email format"`
	Password  string     `gorm:"not null" json:"password" form:"password" `
	Age       uint       `gorm:"not null" json:"age" form:"age" `
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAT *time.Time `json:"updated_at,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
