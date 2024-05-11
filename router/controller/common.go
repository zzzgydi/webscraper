package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/webscraper/router/utils"
)

// Code
const (
	CodeSuccess     int = 0
	CodeBadRequest  int = 400
	CodeServerError int = 500
)

// Msg
const (
	MsgSuccess = "ok"
	MsgFailed  = "failed"
)

type CommonResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func returnResponse(c *gin.Context, code int, msg string, data any) {
	if trace := utils.GetTraceLogger(c); trace != nil {
		trace.BizCode = code
		trace.BizMsg = msg
	}

	c.JSON(http.StatusOK, &CommonResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func ReturnSuccess(c *gin.Context, data any) {
	returnResponse(c, CodeSuccess, MsgSuccess, data)
}

func ReturnServerError(c *gin.Context, err error) {
	returnResponse(c, CodeServerError, err.Error(), nil)
}

func ReturnBadRequest(c *gin.Context, err error) {
	returnResponse(c, CodeBadRequest, err.Error(), nil)
}
