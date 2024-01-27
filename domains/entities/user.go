package entities

import "time"

type User struct {
	ID        int       `gorm:"primaryKey"`
	Email     string    `gorm:"size:50;unique"`
	Password  string    `gorm:"size:100;"`
	FullName  string    `gorm:"size:40;"`
	Address   string    `gorm:"size:256;"`
	Age       int32     `gorm:"column:age"`
	Sex       string    `gorm:"size:10;"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
