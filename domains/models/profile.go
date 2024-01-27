package models

import "time"

type ListProfileResponse struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Age      int32  `json:"age"`
	Sex      string `json:"sex"`
}

type LoginUserPremium struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	ExpiredAt time.Time `json:"expired_at"`
}
