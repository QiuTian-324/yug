package libs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standardized response structure for all API responses.
type Response struct {
	Code    int         `json:"code"`            // 业务状态码：0 为成功，-1 为失败
	Message string      `json:"message"`         // 消息内容
	Success bool        `json:"success"`         // 是否成功
	Data    interface{} `json:"data"`            // 返回数据
	Extra   interface{} `json:"extra,omitempty"` // 可选的扩展数据
}

// 自定义的业务状态码
const (
	CodeSuccess = 0
	CodeFail    = -1
)

// NewResponse 构造通用响应
func NewResponse(code int, message string, success bool, data interface{}, extra interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Success: success,
		Data:    data,
		Extra:   extra,
	}
}

// sendResponse 通过 ctx 返回 JSON 响应，同时设置 HTTP 状态码
func sendResponse(ctx *gin.Context, httpCode int, response Response) {
	ctx.AbortWithStatusJSON(httpCode, response)
}

// OK 返回成功的响应
func OK(ctx *gin.Context, message string, data interface{}) {
	response := NewResponse(CodeSuccess, message, true, data, nil)
	sendResponse(ctx, http.StatusOK, response)
}

// Created 返回资源创建成功的响应
func Created(ctx *gin.Context, message string, data interface{}) {
	response := NewResponse(CodeSuccess, message, true, data, nil)
	sendResponse(ctx, http.StatusCreated, response)
}

// Failed 返回失败的响应
func Failed(ctx *gin.Context, message string, data interface{}) {
	response := NewResponse(CodeFail, message, false, data, nil)
	sendResponse(ctx, http.StatusBadRequest, response)
}

// NotFound 返回未找到的响应
func NotFound(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusNotFound, response)
}

// Unauthorized 返回未授权的响应
func Unauthorized(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusUnauthorized, response)
}

// Forbidden 返回禁止访问的响应
func Forbidden(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusForbidden, response)
}

// TooMany 返回请求过于频繁的响应
func TooMany(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusTooManyRequests, response)
}

// ParamError 返回参数错误的响应
func ParamError(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusBadRequest, response)
}

// Internal 返回服务器错误的响应
func Internal(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusInternalServerError, response)
}

// ServiceUnavailable 返回服务不可用的响应
func ServiceUnavailable(ctx *gin.Context, message string) {
	response := NewResponse(CodeFail, message, false, nil, nil)
	sendResponse(ctx, http.StatusServiceUnavailable, response)
}

// AddExtra 可以向现有响应添加额外信息
func AddExtra(ctx *gin.Context, response Response, extra interface{}) {
	response.Extra = extra
	sendResponse(ctx, http.StatusOK, response)
}
