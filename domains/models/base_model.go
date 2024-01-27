package models

type BaseRequest struct {
	ID         int               `mapper:"id"`
	Page       int               `mapper:"page"`
	Limit      int               `mapper:"limit"`
	Token      string            `mapper:"token"`
	BodyData   interface{}       `mapper:"data"`
	QueryParam map[string]string `mapper:"query_param"`
	Param      uint              `mapper:"param"`
}
