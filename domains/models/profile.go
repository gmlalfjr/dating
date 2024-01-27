package models

type ListProfileResponse struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Age      int32  `json:"age"`
	Sex      string `json:"sex"`
}
