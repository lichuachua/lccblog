package syserror

type HasPraiseError struct {
	UnKnowError
}
func (this HasPraiseError) Code() int {
	return 4444
}
func (this HasPraiseError) Error() string {
	return "已经点赞过"
}