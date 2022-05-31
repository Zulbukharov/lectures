package healthchecker

import "errors"

var (
	ErrFailed = errors.New("website is not available")
)
