package errors

import (
	"fmt"
)

type DataNotFoundError struct {
	resource string
	id       string
}

func (r DataNotFoundError) Error() string {
	return r.Message()
}

func (r DataNotFoundError) Message() string {
	return fmt.Sprintf("DataNotFound: %v", r.id)
}

func (r DataNotFoundError) Code() string {
	return "DataNotFound"
}

func NewDataNotFoundError(resource string, id string) DataNotFoundError {
	return DataNotFoundError{
		id:       id,
		resource: resource,
	}
}
