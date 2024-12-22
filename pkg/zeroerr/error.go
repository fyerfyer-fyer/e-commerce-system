package zeroerr

import "fmt"

type ErrorCode int

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

const (
	ErrInternal     ErrorCode = iota + 1000 // 内部错误
	ErrParam                                // 参数错误
	ErrNotFound                             // 资源未找到
	ErrUnauthorized                         // 未授权
	ErrForbidden                            // 禁止访问
	// ... 其他错误码
)

type MyError struct {
	Code    ErrorCode
	Message string
}

var geneMsg map[ErrorCode]string

// 实现 Error 方法，使 MyError 类型符合 error 接口
func (e *MyError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func NewError(code ErrorCode, message string) *MyError {
	return &MyError{
		Code:    code,
		Message: message,
	}
}

func init() {
	geneMsg = make(map[ErrorCode]string)

	geneMsg[ErrInternal] = "内部错误"
	geneMsg[ErrParam] = "参数错误"
	geneMsg[ErrNotFound] = "资源未找到"
	geneMsg[ErrUnauthorized] = "未授权"
	geneMsg[ErrForbidden] = "禁止访问"
}

func GetErrorMessage(code ErrorCode) string {
	if msg, exists := geneMsg[code]; exists {
		return msg
	}
	return "未知错误"
}
