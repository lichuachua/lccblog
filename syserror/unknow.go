package syserror

type UnKnowError struct {

	msg string
	reason error
}

func (this UnKnowError) Code() int {
	return 1000
}
func (this UnKnowError) Error() string {
	if len(this.msg)==0 {
		return "未知错误"
	}
	return this.msg
}
func (this UnKnowError) ReasonError() error {
	return this.reason
}
