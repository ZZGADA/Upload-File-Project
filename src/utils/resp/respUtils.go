package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResultCont struct {
	Code int32       `json:"code" mapstructure:"code"`
	Msg  string      `json:"msg" mapstructure:"msg"`
	Data interface{} `json:"data" mapstructure:"data"`
}

type Result struct {
	Ctx *gin.Context
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{
		Ctx: ctx,
	}
}

func NewResultCont(code int32, msg string, data interface{}) *ResultCont {
	return &ResultCont{Code: code, Msg: msg, Data: data}
}

func NewResultContEmpty() *ResultCont {
	return &ResultCont{}
}

// Success
func (result *Result) Success(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := NewResultContEmpty()
	res.Code = 200
	res.Msg = ""
	res.Data = data
	result.Ctx.JSON(http.StatusOK, res)
}

// Success
func (result *Result) Failed(statusCode int, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := NewResultContEmpty()
	res.Code = int32(statusCode)
	res.Msg = ""
	res.Data = data
	result.Ctx.JSON(statusCode, res)
}
