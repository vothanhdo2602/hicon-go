package errorcode

import "errors"

type CommonErrorCode int
type CommonErrors map[CommonErrorCode]string

func GetCommonError(code CommonErrorCode) error {
	return errors.New(commons[code])
}

const (
	CommonNotFound = "common_not_found"
)

const (
	CommonNotFoundCode CommonErrorCode = iota
)

var commons = CommonErrors{
	CommonNotFoundCode: CommonNotFound,
}
