package randutils

import (
	"crypto/rand"
	"errors"
	"math/big"
)

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
	result := make([]byte, length)
	if setLength > 256 {
		for i := range result {
			indexOfSet, err := rand.Int(r, big.NewInt(int64(setLength)))
			if err != nil {
				return RandomGenerator{}, err
			}
			result[i] = set[indexOfSet.Int64()]
		}
		return RandomGenerator{
			Data: result,
		}, nil
	}
	maxByte := byte(255 - (255 % setLength))
	_, err := r.ReadRange(result, [2]byte{0, maxByte})
	if err != nil {
		return RandomGenerator{}, err
	}
	for i := range result {
		index := int(result[i]) % setLength
		result[i] = set[index]
	}
	return RandomGenerator{
		Data: result,
	}, nil
}
