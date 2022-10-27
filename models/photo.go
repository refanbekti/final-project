// package models

// import (
// 	"time"

// 	"github.com/asaskevich/govalidator"
// 	"gorm.io/gorm"
// )

// type Photo struct {
// 	GormModel
// 	ID        uint   `gorm:"primaryKey" json:"id"`
// 	Title     string `gorm:"not null" json:"title" form:"full_name"`
// 	Caption   string `gorm:"not null" json:"caption" form:"email"`
// 	Photo_url string `gorm:"not null" json:"photo_url" form:"password" `
// 	UserID    uint
// 	CreatedAt *time.Time `json:"created_at,omitempty"`
// 	UpdatedAT *time.Time `json:"updated_at,omitempty"`
// }

// func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
// 	_, errCreate := govalidator.ValidateStruct(p)
// 	if errCreate != nil {
// 		err = errCreate
// 		return
// 	}

// 	err = nil
// 	return
// }

// func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
// 	_, errCreate := govalidator.ValidateStruct(p)
// 	if errCreate != nil {
// 		err = errCreate
// 		return
// 	}

// 	err = nil
// 	return
// }