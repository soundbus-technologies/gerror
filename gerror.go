// authors: wangoo
// created: 2018-07-04
// golang error extension

package gerror

import "encoding/json"

type CodeError interface {
	error
	Code() string
	Msg() string
}

type serror struct {
	errCode string
	errMsg  string
}

func (e *serror) Code() string {
	return e.errCode
}

func (e *serror) Msg() string {
	return e.errMsg
}

func (e *serror) Error() string {
	return e.errCode + ":" + e.errMsg
}

type ServiceError struct {
	Status      int    `json:"status"`
	ErrCode     string `json:"errCode"`
	ErrMsg      string `json:"errMsg,omitempty"`
	DevMsg      string `json:"devMsg,omitempty"`
	MoreInfoUrl string `json:"moreInfoUrl,omitempty"`
	RequestId   string `json:"requestId,omitempty"`
}

func (e *ServiceError) Code() string {
	return e.ErrCode
}

func (e *ServiceError) Msg() string {
	return e.ErrMsg
}

func (e *ServiceError) Error() string {
	if b, err := json.Marshal(e); err == nil {
		return string(b)
	}
	return e.ErrCode + ":" + e.ErrMsg
}

func New(code, error string) *serror {
	return &serror{
		errCode: code,
		errMsg:  error,
	}
}

func NewServiceError(status int, errCode, errMsg, devMsg, infoUrl, requestId string) *ServiceError {
	return &ServiceError{
		Status:      status,
		ErrCode:     errCode,
		ErrMsg:      errMsg,
		DevMsg:      devMsg,
		MoreInfoUrl: infoUrl,
		RequestId:   requestId,
	}
}

func NewServiceCodeError(codeErr CodeError, status int, devMsg string) *ServiceError {
	return NewServiceDetailError(codeErr, status, devMsg, "", "")
}

func NewServiceDetailError(codeErr CodeError, status int, devMsg, infoUrl, requestId string) *ServiceError {
	return NewServiceError(status, codeErr.Code(), codeErr.Msg(), devMsg, infoUrl, requestId)
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
