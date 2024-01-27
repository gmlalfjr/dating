package response

import (
	"dating/constants"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IResponse interface {
	Success(c *gin.Context)
}

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	TotalPages int         `json:"total_page,omitempty"`
	TotalData  int         `json:"total_data,omitempty"`
}

func (res *Response) Success(c *gin.Context) {
	if res.StatusCode == 0 {
		res.StatusCode = http.StatusOK
	}

	res.Message = "OK"

	c.JSON(http.StatusOK, res)
}

func (res *Response) Error(c *gin.Context, err error) {
	var resp = &Response{
		StatusCode: http.StatusInternalServerError,
		Message:    constants.InternalServerError,
	}

	if err != nil {
		switch er := err.(type) {
		case *Error:
			resp = &Response{
				StatusCode: er.statusCode,
				Message:    er.message,
			}
		}
	}
	c.JSON(resp.StatusCode, resp)

}
