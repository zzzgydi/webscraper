package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zzzgydi/webscraper/router/utils"
)

const (
	SUCCESS      int = 0
	BAD_REQUEST  int = 400
	SERVER_ERROR int = 500
)

type ObjectResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func ReturnSuccess(c *gin.Context, data any) {
	if trace := utils.GetTraceLogger(c); trace != nil {
		trace.BizCode = SUCCESS
		trace.BizMsg = "ok"
	}

	c.JSON(http.StatusOK, &ObjectResponse{
		Code: SUCCESS,
		Msg:  "ok",
		Data: data,
	})
}

func ReturnServerError(c *gin.Context, err error) {
	if trace := utils.GetTraceLogger(c); trace != nil {
		trace.BizCode = SERVER_ERROR
		trace.BizMsg = err.Error()
	}

	c.JSON(http.StatusOK, &ObjectResponse{
		Code: SERVER_ERROR,
		Msg:  err.Error(),
	})
}

func ReturnBadRequest(c *gin.Context, err error) {
	if trace := utils.GetTraceLogger(c); trace != nil {
		trace.BizCode = BAD_REQUEST
		trace.BizMsg = err.Error()
	}

	c.JSON(http.StatusOK, &ObjectResponse{
		Code: BAD_REQUEST,
		Msg:  err.Error(),
	})
}
