package libs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Resp struct {
	statusCode int
	errCode    int64
	msg        interface{}
	data       interface{}
	debug      interface{}

	ctx    *gin.Context
	isProd bool
}

const (
	ErrCodeOk = 0 // 没有错误
)

func NewResp(ctx *gin.Context, isProd bool) (r *Resp) {
	r = &Resp{
		statusCode: http.StatusOK,
		errCode:    ErrCodeOk,
	}
	r.isProd = isProd
	r.ctx = ctx
	return r
}

// StatusCode 设置响应的http状态码
func (r *Resp) StatusCode(code int) *Resp {
	r.statusCode = code
	return r
}

// ErrCode 设置业务自定义错误码,默认为0
func (r *Resp) ErrCode(errCode int64) *Resp {
	r.errCode = errCode
	return r
}

// Data 设置响应的主要数据data
func (r *Resp) Data(data interface{}) *Resp {
	r.data = data
	return r
}

// Msg 设置响应的提示信息或错误信息
func (r *Resp) Msg(msg interface{}) *Resp {
	r.msg = msg
	return r
}

// Debug 设置响应的debug信息,仅非生产模式会输出
func (r *Resp) Debug(debugInfo interface{}) *Resp {
	if r.isProd {
		return r
	}
	r.debug = debugInfo
	return r
}

// Echo 输出响应
func (r *Resp) Echo() {
	ctx := r.ctx
	ctx.Status(r.statusCode)

	if r.statusCode != http.StatusOK &&
		r.statusCode != http.StatusBadRequest {
		return
	}

	s := struct {
		ErrCode int64       `json:"errCode"`
		Data    interface{} `json:"data,omitempty"`
		Msg     interface{} `json:"msg,omitempty"`
		Debug   interface{} `json:"debug,omitempty"`
	}{
		ErrCode: r.errCode,
		Data:    r.data,
		Msg:     r.msg,
		Debug:   r.debug,
	}
	ctx.JSON(r.statusCode, s)
}

// 200 状态码为200的 Echo 快捷输出
func (r *Resp) Ok() {
	r.StatusCode(http.StatusOK).Echo()
}

// 400 状态码为400的 Echo 快捷输出
func (r *Resp) Err(errMsg interface{}) {
	r.Msg(errMsg).StatusCode(http.StatusBadRequest).Echo()
}
