package response

import "strings"

type Error struct {
	err        error
	message    string
	moreInfo   string
	statusCode int
	systemCode int
}

// NewError returns a new strandard error
func NewError(err error, statusCode int, message string, moreInfos ...string) error {
	return &Error{
		err:        err,
		statusCode: statusCode,
		message:    message,
		moreInfo:   strings.Join(moreInfos, ","),
	}
}

func (e *Error) Error() string {
	return e.err.Error()

}
