package syserror

type UnknownError struct {
	msg    string
	reason error
}

func (this UnknownError) Code() int {
	return 1000
}
func (this UnknownError) Error() string {
	if len(this.msg) == 0 {
		return "unknown error"
	}
	return this.msg
}
func (this UnknownError) ReasonError() error {
	return this.reason
}
