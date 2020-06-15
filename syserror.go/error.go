package syserror

//定义错误类型
type Error interface {
	Code() int
	Error() string
	ReasonError() error
}

func New(msg string, reason error) Error {
	return UnknownError{msg: msg, reason: reason}
}
