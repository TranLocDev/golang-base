package models

import (
	"time"
)

// Students định nghĩa cấu trúc dữ liệu cho bảng students
type Students struct {
	Id            uint      `json:"id" gorm:"primarykey;autoIncrement"`
	FirstName     *string   `json:"first_name" gorm:"not null"`
	LastName      *string   `json:"last_name" gorm:"not null"`
	Email         *string   `json:"email" gorm:"not null"`
	Phone         *string   `json:"phone" gorm:"not null"`
	Country       *string   `json:"country" gorm:"not null"`
	CountryCode   *string   `json:"country_code" gorm:"not null"`
	CountryID     uint      `json:"country_id" gorm:"not null"`
	CountryDetail Countries `json:"country_detail" gorm:"foreignKey:CountryID;references:Id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// TableName đặt tên bảng tương ứng
func (Students) TableName() string {
	return "students"
}

