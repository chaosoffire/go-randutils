package randutils

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
)

// precision for big.Float operations
const (
	precision = 256
)

var (
	Reader *RandBufferReader

	numset    = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	lowerset  = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	upperset  = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	symbolset = []byte{'!', '@', '#', '$', '%', '^', '&', '*'}

	alphabetset = append(upperset, lowerset...)
	charset     = append(alphabetset, numset...)
	allset      = append(charset, symbolset...)
)

func init() {
	var err error
	Reader, err = NewRandBufferReader()
	if err != nil {
		panic(fmt.Sprintf("randutils: failed to initialize cryptographic random source: %v", err))
	}
}

type RandomGenerator struct {
	Data any
}

func (r *RandomGenerator) String() (string, error) {
	switch v := r.Data.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case int:
		return fmt.Sprintf("%d", v), nil
	case int32:
		return fmt.Sprintf("%d", v), nil
	case int64:
		return fmt.Sprintf("%d", v), nil
	case *big.Int:
		return v.String(), nil
	case float32:
		return fmt.Sprintf("%f", v), nil
	case float64:
		return fmt.Sprintf("%f", v), nil
	case *big.Float:
		return v.Text('f', precision), nil
	default:
		return "", errors.New("Data is not a string")
	}
}

func (r *RandomGenerator) Int() (int64, error) {
	switch v := r.Data.(type) {
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case *big.Int:
		if v.IsInt64() {
			return v.Int64(), nil
		}
		return 0, errors.New("value is too large to fit in int")
	default:
		return 0, errors.New("data cannot be converted to an int")
	}
}

func (r *RandomGenerator) BigInt() (*big.Int, error) {
	switch v := r.Data.(type) {
	case int:
		return new(big.Int).SetInt64(int64(v)), nil
	case int32:
		return new(big.Int).SetInt64(int64(v)), nil
	case int64:
		return new(big.Int).SetInt64(v), nil
	case *big.Int:
		return v, nil
	default:
		return nil, errors.New("data cannot be converted to a BigInt")
	}
}

func (r *RandomGenerator) Float() (float64, error) {
	switch v := r.Data.(type) {
	case int:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case *big.Float:
		if f, _ := v.Float64(); f != math.Inf(1) && f != math.Inf(-1) {
			return f, nil
		} else {
			return 0, errors.New("value is too large or too small to fit in float64")
		}
	default:
		return 0, errors.New("data cannot be converted to a float")
	}
}

func (r *RandomGenerator) BigFloat() (*big.Float, error) {
	switch v := r.Data.(type) {
	case int:
		return new(big.Float).SetPrec(precision).SetInt64(int64(v)), nil
	case int32:
		return new(big.Float).SetPrec(precision).SetInt64(int64(v)), nil
	case int64:
		return new(big.Float).SetPrec(precision).SetInt64(v), nil
	case float32:
		return new(big.Float).SetPrec(precision).SetFloat64(float64(v)), nil
	case float64:
		return new(big.Float).SetPrec(precision).SetFloat64(v), nil
	case *big.Float:
		return v, nil
	default:
		return nil, errors.New("data cannot be converted to a BigFloat")
	}
}

func (r *RandomGenerator) Bytes() ([]byte, error) {
	switch v := r.Data.(type) {
	case []byte:
		return v, nil
	default:
		return nil, errors.New("data is not bytes")
	}
}

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
