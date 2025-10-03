package randutils

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

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
		return v.Text('f', -1), nil
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
