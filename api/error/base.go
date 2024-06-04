package errs

import (
	"codegen/api/code"
	"codegen/pkg/server/api"
	"net/http"
)

var (

	// Failed 操作失败
	Failed = &Err{
		code:     code.Failed,
		httpCode: 200,
		message:  "操作失败",
	}

	// RecordNotFound 数据不存在
	RecordNotFound = &Err{
		code:     code.Failed,
		httpCode: 200,
		message:  "数据不存在",
	}

	// ParamErr 参数错误
	ParamErr = &Err{
		code:     code.ParamErr,
		httpCode: 200,
		message:  "参数错误",
	}

	// RecordExisting 数据已存在
	RecordExisting = &Err{
		code:     code.Failed,
		httpCode: 200,
		message:  "数据已存在",
	}

	// RecordExisting 数据已存在
	DatasourceConnFailed = &Err{
		code:     code.Failed,
		httpCode: 200,
		message:  "数据源链接失败",
	}
)

// Err 错误类
type Err struct {
	code     int
	httpCode int
	message  string
}

// Code 错误码
func (err *Err) Code() int {
	return err.code
}

// HttpCode Http状态码
func (err *Err) HttpCode() int {
	return err.httpCode
}

// Error 错误信息
func (err *Err) Error() string {
	return err.message
}

// ErrResp 操作失败响应
func ErrResp(err api.Error) (httpCode int, rsp Response) {
	httpCode = err.HttpCode()
	rsp = Response{
		Code: err.Code(),
		Msg:  err.Error(),
	}
	return
}

// Response 响应基本数据
type Response struct {
	Code int         `json:"code" validate:"required"` // 响应码
	Msg  string      `json:"msg" validate:"required"`  // 响应消息
	Data interface{} `json:"data"`                     // 响应数据
}

// SucResp 操作成功响应
func SucResp(data interface{}) (resCode int, res Response) {
	resCode = 200
	res = Response{
		Code: code.Success,
		Msg:  "操作成功",
		Data: data,
	}
	return
}

// NewErrorResp 生成错误响应 message code httpCode
func NewErrorResp(message string, args ...int) (int, Response) {
	var (
		badRequestCode = http.StatusBadRequest
		httpCode       = http.StatusOK
	)

	for key, arg := range args {
		if key == 0 {
			badRequestCode = arg
		} else if key == 1 {
			httpCode = arg
		}
	}
	return ErrResp(&Err{code: badRequestCode, httpCode: httpCode, message: message})
}
