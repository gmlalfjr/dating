package models

type BaseRequest struct {
	ID         int               `mapper:"id"`
	Page       int               `mapper:"page"`
	Limit      int               `mapper:"limit"`
	Token      string            `mapper:"token"`
	BodyData   interface{}       `mapper:"data"`
	QueryParam map[string]string `mapper:"query_param"`
	Param      uint              `mapper:"param"`
	User       UserRequest       `mapper:"user"`
}

type UserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
