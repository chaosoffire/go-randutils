package randutils

import (
	"errors"
	"io"
)

func NewBytes(r *RandBufferReader, length int) (RandomGenerator, error) {
	if length <= 0 {
		return RandomGenerator{}, errors.New("length of bytes slice must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	result := make([]byte, length)
	_, err := io.ReadFull(r, result)
	if err != nil {
		return RandomGenerator{}, err
	}
	return RandomGenerator{
		Data: result,
	}, nil
}
