package entities

import "time"

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Email     string    `gorm:"size:50;unique" json:"email"`
	Password  string    `gorm:"size:100;" json:"size:100;password"`
	FullName  string    `gorm:"size:40;" json:"full_name"`
	Address   string    `gorm:"size:256;" json:"address"`
	Age       int32     `json:"age"`
	Sex       string    `gorm:"size:10;" json:"sex"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}
