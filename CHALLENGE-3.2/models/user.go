package models

import (
	"fmt"
	"regexp"

	"CHALLENGE-3.2/helpers"
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"type:varchar(100);unique_index" validate:"required,email"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	IsAdmin  bool      `gorm:"default:0" json:"is_admin"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:sSET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(u.Email) {
		return fmt.Errorf("invalid email format")
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
