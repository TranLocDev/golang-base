package models

import (
	"time"
)

// Countries định nghĩa cấu trúc dữ liệu cho bảng countries
type Countries struct {
	Id          uint      `json:"id" gorm:"primarykey;autoIncrement"`
	Country     *string   `json:"country" gorm:"not null"`
	CountryCode *string   `json:"countryCode" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName đặt tên bảng tương ứng
func (Countries) TableName() string {
	return "countries"
}

