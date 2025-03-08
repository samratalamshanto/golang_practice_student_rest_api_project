package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	MSG     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Err     interface{} `json:"err"`
}

func SuccessResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, CommonResponse{
		Code:    http.StatusOK,
		Success: true,
		Message: msg,
		MSG:     msg,
		Data:    data, //Go internally resolves the pointer before converting it to JSON. Automatically handles dereferencing when it marshals the data into JSON.
		Error:   nil,
	})
}

func ErrorResponse(c *gin.Context, msg string, err interface{}) {
	c.JSON(http.StatusInternalServerError, CommonResponse{
		Code:    http.StatusInternalServerError,
		Success: false,
		Message: msg,
		MSG:     msg,
		Error:   err,
		Err:     err,
		Data:    nil,
	})
}
