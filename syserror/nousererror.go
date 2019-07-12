package syserror

type NoUserError struct {
	UnKnowError
}

func (this NoUserError) Code() int {
	return 1001
}
func (this NoUserError) Error() string {
		return "请登录博客"
}