package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IResponse interface {
	Success(c *gin.Context)
}

type Response struct {
	Status     int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	TotalPages int         `json:"total_page,omitempty"`
	TotalData  int         `json:"total_data,omitempty"`
}

func (res *Response) Success(c *gin.Context) {
	if res.Status == 0 {
		res.Status = http.StatusOK
	}

	res.Message = "OK"

	c.JSON(http.StatusOK, res)
}

func (res Response) Error(c *gin.Context, code int, message ...string) {
	res.Status = code
	res.Message = message[0]
	c.JSON(code, res)
}
