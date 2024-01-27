package entities

import "time"

type Swipe struct {
	ID           int `gorm:"primaryKey"`
	UserID       int
	User         User `gorm:"foreignKey:UserID"`
	SwipedUserID int
	SwipedUserId User      `gorm:"foreignKey:SwipedUserID"`
	Action       string    `gorm:"size:10;" json:"action"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
}

//SwipedUserId int       `json:"swiped_user_id"`
//SwipedUser   *User     `gorm:"foreignKey:SwipedUserId" `
