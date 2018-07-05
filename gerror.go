// authors: wangoo
// created: 2018-07-04
// golang error extension

package gerror

import (
	"fmt"
	"encoding/json"
)

type ErrorCoder interface {
	error
	Code() string
	Json() string
}

type codeError struct {
	code  string
	error string
}

func (e *codeError) Code() string {
	return e.code
}

func (e *codeError) Error() string {
	return e.error
}

func (e *codeError) Json() string {
	c, _ := json.Marshal(e.code)
	r, _ := json.Marshal(e.error)
	return fmt.Sprintf("{\"code\":%v,\"error\":%v}", string(c), string(r))
}

func New(code, error string) *codeError {
	return &codeError{
		code:  code,
		error: error,
	}
}

// see http://docs.oifitech.com/pages/viewpage.action?pageId=3539165
var (
	ErrInternalError          = New("E0000", "系统内部错误")
	ErrServiceNotAvailable    = New("E0001", "服务不可用 Service Not Available")
	ErrSystemIdRequired       = New("E0002", "无法获取SystemId")
	ErrAssertError            = New("E0003", "流程中断异常")
	ErrConditionDissatisfied  = New("E0004", "前置条件不满足")
	ErrDuplicated             = New("E0005", "资源已经存在")
	ErrOperationUnsupported   = New("E0006", "不支持的操作")
	ErrHttpMethodUnsupported  = New("E1100", "HTTP方法错误")
	ErrContentTypeUnsupported = New("E1101", "HTTP Header Content-Type 错误")
	ErrTimeout                = New("E1102", "请求超时错误")
	ErrFrequently             = New("E1103", "请求过于频繁")
	ErrIPForbidden            = New("E1104", "禁止IP地址访问")
	ErrValueRequired          = New("E2000", "必填项检查")
	ErrUniqueConstraint       = New("E2001", "唯一性检查错误")
	ErrValueInvalid           = New("E2002", "有效性检查错误")
	ErrValueUnsupported       = New("E2003", "无效值检查错误")
	ErrTypeInvalid            = New("E2004", "类型检查错误")
	ErrNotFound               = New("E2005", "无法找到对象")
)
