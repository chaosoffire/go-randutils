package randutils

import "errors"

func NewRandomFromSet(r *RandBufferReader, length int, set []byte) (RandomGenerator, error) {
	if length <= 0 {
		return RandomGenerator{}, errors.New("length must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	setLength := len(set)
	if setLength == 0 {
		return RandomGenerator{}, errors.New("given set must not be empty")
	}

	maxByte := byte(255 - (255 % setLength))
	result := make([]byte, length)
	_, err := r.ReadRange(result, [2]byte{0, maxByte})
	if err != nil {
		return RandomGenerator{}, err
	}
	for i := range result {
		result[i] = set[result[i]%byte(setLength)]
	}
	return RandomGenerator{
		Data: result,
	}, nil
}
