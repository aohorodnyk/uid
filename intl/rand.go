package intl

import (
	"errors"
	"io"
)

const ErrorMsgWrongSizeZero = "size cannot be zero"

type Rand struct {
	reader io.Reader
}

func (r Rand) Generate(size uint32) ([]byte, error) {
	if size == 0 {
		return nil, &WrongSizeError{errors.New(ErrorMsgWrongSizeZero)}
	}

	b := make([]byte, size)
	_, err := r.reader.Read(b)
	if err != nil {
		return nil, &InternalError{err}
	}

	return b, nil
}
