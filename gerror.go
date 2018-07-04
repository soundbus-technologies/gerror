// authors: wangoo
// created: 2018-07-04
// golang error extension

package gerror

type ErrorCoder interface {
	error
	Code() string
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

func NewCodeError(code, error string) *codeError {
	return &codeError{
		code:  code,
		error: error,
	}
}

// see http://docs.oifitech.com/pages/viewpage.action?pageId=3539165
var (
	ErrInternalError          = NewCodeError("E0000", "系统内部错误")
	ErrServiceNotAvailable    = NewCodeError("E0001", "服务不可用 Service Not Available")
	ErrSystemIdRequired       = NewCodeError("E0002", "无法获取SystemId")
	ErrAssertError            = NewCodeError("E0003", "流程中断异常")
	ErrConditionDissatisfied  = NewCodeError("E0004", "前置条件不满足")
	ErrDuplicated             = NewCodeError("E0005", "资源已经存在")
	ErrOperationUnsupported   = NewCodeError("E0006", "不支持的操作")
	ErrHttpMethodUnsupported  = NewCodeError("E1100", "HTTP方法错误")
	ErrContentTypeUnsupported = NewCodeError("E1101", "HTTP Header Content-Type 错误")
	ErrTimeout                = NewCodeError("E1102", "请求超时错误")
	ErrFrequently             = NewCodeError("E1103", "请求过于频繁")
	ErrIPForbidden            = NewCodeError("E1104", " 禁止IP地址访问")
	ErrValueRequired          = NewCodeError("E2000", "必填项检查")
	ErrUniqueConstraint       = NewCodeError("E2001", "唯一性检查错误")
	ErrValueInvalid           = NewCodeError("E2002", "有效性检查错误")
	ErrValueUnsupported       = NewCodeError("E2003", "无效值检查错误")
	ErrTypeInvalid            = NewCodeError("E2004", "类型检查错误")
	ErrNotFound               = NewCodeError("E2005", "无法找到对象")
)
