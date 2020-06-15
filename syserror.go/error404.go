package syserror

type Error404 struct {
	UnknownError
}

func (this Error404) Code() int {
	return 404
}
func (this Error404) Error() string {
	return "page not found"
}
func (this Error404) ReasonError() error {
	return this.reason
}
