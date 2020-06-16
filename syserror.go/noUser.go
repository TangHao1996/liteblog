package syserror

type ErrorNoUser struct {
	UnknownError
}

func (this ErrorNoUser) Code() int {
	return 1001
}
func (this ErrorNoUser) Error() string {
	return "请先登录！"
}
