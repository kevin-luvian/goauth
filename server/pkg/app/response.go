package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevin-luvian/goauth/server/pkg/e"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(e.SUCCESS),
		Data: data,
	})
}

func Error(c *gin.Context, code int, err error) {
	c.JSON(e.GetStatus(code), Response{
		Code: e.SUCCESS,
		Msg:  e.GetMsg(code),
		Data: err.Error(),
	})
}
