package randutils

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"math"
	"math/big"
)

func NewInt(r *RandBufferReader, maximum int) (RandomGenerator, error) {
	if maximum <= 0 {
		return RandomGenerator{}, errors.New("maximum number must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	result, err := rand.Int(r, big.NewInt(int64(maximum)))
	if err != nil {
		return RandomGenerator{}, err
	}
	return RandomGenerator{
		Data: int(result.Int64()),
	}, nil
}

func NewBigInt(r *RandBufferReader, maximum *big.Int) (RandomGenerator, error) {
	if maximum.Cmp(big.NewInt(0)) <= 0 {
		return RandomGenerator{}, errors.New("maximum number must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	result, err := rand.Int(r, maximum)
	if err != nil {
		return RandomGenerator{}, err
	}
	return RandomGenerator{
		Data: result,
	}, nil
}

func NewBigIntLength(r *RandBufferReader, length int) (RandomGenerator, error) {
	if length <= 0 {
		return RandomGenerator{}, errors.New("length must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	min := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length-1)), nil)
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)
	rangeSize := new(big.Int).Sub(max, min)
	result, err := rand.Int(r, rangeSize)
	if err != nil {
		return RandomGenerator{}, err
	}
	result = new(big.Int).Add(result, min)
	return RandomGenerator{
		Data: result,
	}, nil
}

func NewFloat(r *RandBufferReader, maximum float64) (RandomGenerator, error) {
	if maximum <= 0 {
		return RandomGenerator{}, errors.New("maximum number must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	var randomUint64 uint64
	err := binary.Read(r, binary.BigEndian, &randomUint64)
	if err != nil {
		return RandomGenerator{}, err
	}
	f01 := math.Float64frombits(0x3FF0000000000000|(randomUint64>>12)) - 1.0
	result := f01 * maximum
	return RandomGenerator{
		Data: result,
	}, nil
}

func NewBigFloat(r *RandBufferReader, max *big.Float) (RandomGenerator, error) {
	if max.Cmp(big.NewFloat(0)) <= 0 {
		return RandomGenerator{}, errors.New("maximum number must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	// maxInt = 2^256
	maxInt := new(big.Int).Lsh(big.NewInt(1), precision)
	randInt, err := rand.Int(r, maxInt)
	if err != nil {
		return RandomGenerator{}, err
	}
	randFloat := new(big.Float).SetPrec(precision).SetInt(randInt)
	maxFloat := new(big.Float).SetPrec(precision).SetInt(maxInt)
	f01 := new(big.Float).SetPrec(precision).Quo(randFloat, maxFloat)
	result := new(big.Float).SetPrec(precision).Mul(f01, max)
	return RandomGenerator{
		Data: result,
	}, nil
}

func NewBigFloatLength(r *RandBufferReader, length int) (RandomGenerator, error) {
	if length <= 0 {
		return RandomGenerator{}, errors.New("length must be greater than 0")
	}
	if r == nil {
		r = Reader
	}
	integerPart, err := NewBigIntLength(r, length)
	if err != nil {
		return RandomGenerator{}, err
	}
	decimalPart, err := NewBigFloat(r, new(big.Float).SetPrec(precision).SetFloat64(1))
	if err != nil {
		return RandomGenerator{}, err
	}
	intPart, err := integerPart.BigInt()
	if err != nil {
		return RandomGenerator{}, err
	}
	decPart, err := decimalPart.BigFloat()
	if err != nil {
		return RandomGenerator{}, err
	}
	result := new(big.Float).
		SetPrec(precision).
		Add(
			new(big.Float).
				SetPrec(precision).
				SetInt(intPart),
			new(big.Float).
				SetPrec(precision).
				Set(decPart))
	return RandomGenerator{
		Data: result,
	}, nil
}
