package entities

import "time"

type PremiumUser struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	User      User      `gorm:"foreignKey:UserID"`
	ExpiredAt time.Time `gorm:"column:expired_at"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}
