package randutils

import "errors"

func NewLowerChars(r *RandBufferReader, length int) (RandomGenerator, error) {
	if length <= 0 {
		return RandomGenerator{}, errors.New("length must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	result, err := NewRandomFromSet(r, length, lowerset)
	if err != nil {
		return RandomGenerator{}, err
	}
	result.Data = string(result.Data.([]byte))
	return result, nil
}

func NewUpperChars(r *RandBufferReader, length int) (RandomGenerator, error) {
	if length <= 0 {
		return RandomGenerator{}, errors.New("length must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	result, err := NewRandomFromSet(r, length, upperset)
	if err != nil {
		return RandomGenerator{}, err
	}
	result.Data = string(result.Data.([]byte))
	return result, nil
}

func NewSymbolChars(r *RandBufferReader, length int) (RandomGenerator, error) {
	if length <= 0 {
		return RandomGenerator{}, errors.New("length must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	result, err := NewRandomFromSet(r, length, symbolset)
	if err != nil {
		return RandomGenerator{}, err
	}
	result.Data = string(result.Data.([]byte))
	return result, nil
}

func NewAlphabets(r *RandBufferReader, length int) (RandomGenerator, error) {
	if length <= 0 {
		return RandomGenerator{}, errors.New("length must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	result, err := NewRandomFromSet(r, length, alphabetset)
	if err != nil {
		return RandomGenerator{}, err
	}
	result.Data = string(result.Data.([]byte))
	return result, nil
}

func NewChars(r *RandBufferReader, length int) (RandomGenerator, error) {
	if length <= 0 {
		return RandomGenerator{}, errors.New("length must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	result, err := NewRandomFromSet(r, length, charset)
	if err != nil {
		return RandomGenerator{}, err
	}
	result.Data = string(result.Data.([]byte))
	return result, nil
}

func NewAllChars(r *RandBufferReader, length int) (RandomGenerator, error) {
	if length <= 0 {
		return RandomGenerator{}, errors.New("length must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	result, err := NewRandomFromSet(r, length, allset)
	if err != nil {
		return RandomGenerator{}, err
	}
	result.Data = string(result.Data.([]byte))
	return result, nil
}
