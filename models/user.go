package models

import (
	"final_project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string        `gorm:"not null;uniqueIndex" json:"username" form:"username" valid:"required~Username is required"`
	Email        string        `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required, email~Invalid email format"`
	Password     string        `gorm:"not null" json:"password" form:"password" valid:"required,minstringlength(6)~Password has to have minimum length of six characters"`
	Age          uint          `gorm:"not null" json:"age" form:"age" valid:"required~Age is required,range(8|100)~Minimum age is 8 years old"`
	Photos       []Photo       `json:"photos" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments     []Comment     `json:"comments" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []SocialMedia `json:"social_medias" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
