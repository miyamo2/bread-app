package errors

type BreadAppError interface {
	Message() string
	Code() string
}
