package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	MSG     string      `json:"msg"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
	Err     interface{} `json:"err"`
}

func SuccessResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, CommonResponse{
		Success: true,
		Message: msg,
		MSG:     msg,
		Data:    data,
		Error:   nil,
	})
}

func ErrorResponse(c *gin.Context, msg string, err interface{}) {
	c.JSON(http.StatusInternalServerError, CommonResponse{
		Success: false,
		Message: msg,
		MSG:     msg,
		Error:   err,
		Err:     err,
		Data:    nil,
	})
}
