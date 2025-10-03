package randutils

import (
	"crypto/rand"
	"errors"
	"io"
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
	// If the set length is greater than 256, we have to use rand.Int for each byte
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
	// If the set length is 256, we can directly read bytes into the result
	if setLength == 256 {
		_, err := io.ReadFull(r, result)
		if err != nil {
			return RandomGenerator{}, err
		}
		for i := range result {
			result[i] = set[i]
		}
		return RandomGenerator{
			Data: result,
		}, nil
	}
	// Standard case for set lengths between 1 and 255
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
