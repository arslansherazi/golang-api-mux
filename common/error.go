package common

type error interface {
	Error() string
}

type GenerateError struct {
	ErrorMessage string
}

func (generateError *GenerateError) Error() string {
	return generateError.ErrorMessage
}
