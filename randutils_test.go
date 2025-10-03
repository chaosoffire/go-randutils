package randutils

import (
	"math/big"
	"testing"
)

func TestNewInt(t *testing.T) {
	max := 100
	rg, err := NewInt(nil, max)
	if err != nil {
		t.Fatalf("NewInt failed: %v", err)
	}
	val, err := rg.Int()
	if err != nil {
		t.Fatalf("Int() failed: %v", err)
	}
	if val < 0 || val >= int64(max) {
		t.Errorf("Int() out of range: got %d, want [0,%d)", val, max)
	}
}

func TestNewBigInt(t *testing.T) {
	max := big.NewInt(1000)
	rg, err := NewBigInt(nil, max)
	if err != nil {
		t.Fatalf("NewBigInt failed: %v", err)
	}
	val, err := rg.BigInt()
	if err != nil {
		t.Fatalf("BigInt() failed: %v", err)
	}
	if val.Cmp(big.NewInt(0)) < 0 || val.Cmp(max) >= 0 {
		t.Errorf("BigInt() out of range: got %v, want [0,%v)", val, max)
	}
}

func TestNewBigIntLength(t *testing.T) {
	length := 5
	rg, err := NewBigIntLength(nil, length)
	if err != nil {
		t.Fatalf("NewBigIntLength failed: %v", err)
	}
	val, err := rg.BigInt()
	if err != nil {
		t.Fatalf("BigInt() failed: %v", err)
	}
	min := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length-1)), nil)
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil)
	if val.Cmp(min) < 0 || val.Cmp(max) >= 0 {
		t.Errorf("BigIntLength() out of range: got %v, want [%v,%v)", val, min, max)
	}
}

func TestNewFloat(t *testing.T) {
	max := 10.0
	rg, err := NewFloat(nil, max)
	if err != nil {
		t.Fatalf("NewFloat failed: %v", err)
	}
	val, err := rg.Float()
	if err != nil {
		t.Fatalf("Float() failed: %v", err)
	}
	if val < 0 || val >= max {
		t.Errorf("Float() out of range: got %f, want [0,%f)", val, max)
	}
}

func TestNewBigFloat(t *testing.T) {
	max := big.NewFloat(100.0)
	rg, err := NewBigFloat(nil, max)
	if err != nil {
		t.Fatalf("NewBigFloat failed: %v", err)
	}
	val, err := rg.BigFloat()
	if err != nil {
		t.Fatalf("BigFloat() failed: %v", err)
	}
	zero := big.NewFloat(0)
	if val.Cmp(zero) < 0 || val.Cmp(max) >= 0 {
		t.Errorf("BigFloat() out of range: got %v, want [0,%v)", val, max)
	}
}

func TestNewBigFloatLength(t *testing.T) {
	length := 4
	rg, err := NewBigFloatLength(nil, length)
	if err != nil {
		t.Fatalf("NewBigFloatLength failed: %v", err)
	}
	val, err := rg.BigFloat()
	if err != nil {
		t.Fatalf("BigFloat() failed: %v", err)
	}
	min := new(big.Float).SetPrec(precision).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length-1)), nil))
	max := new(big.Float).SetPrec(precision).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil))
	if val.Cmp(min) < 0 || val.Cmp(max) >= 0 {
		t.Errorf("BigFloatLength() out of range: got %v, want [%v,%v)", val, min, max)
	}
}

func TestNewLowerChars(t *testing.T) {
	length := 8
	rg, err := NewLowerChars(nil, length)
	if err != nil {
		t.Fatalf("NewLowerChars failed: %v", err)
	}
	str, err := rg.String()
	if err != nil {
		t.Fatalf("String() failed: %v", err)
	}
	if len(str) != length {
		t.Errorf("LowerChars length mismatch: got %d, want %d", len(str), length)
	}
}

func TestNewUpperChars(t *testing.T) {
	length := 8
	rg, err := NewUpperChars(nil, length)
	if err != nil {
		t.Fatalf("NewUpperChars failed: %v", err)
	}
	str, err := rg.String()
	if err != nil {
		t.Fatalf("String() failed: %v", err)
	}
	if len(str) != length {
		t.Errorf("UpperChars length mismatch: got %d, want %d", len(str), length)
	}
}

func TestNewSymbolChars(t *testing.T) {
	length := 8
	rg, err := NewSymbolChars(nil, length)
	if err != nil {
		t.Fatalf("NewSymbolChars failed: %v", err)
	}
	str, err := rg.String()
	if err != nil {
		t.Fatalf("String() failed: %v", err)
	}
	if len(str) != length {
		t.Errorf("SymbolChars length mismatch: got %d, want %d", len(str), length)
	}
}

func TestNewAlphabets(t *testing.T) {
	length := 10
	rg, err := NewAlphabets(nil, length)
	if err != nil {
		t.Fatalf("NewAlphabets failed: %v", err)
	}
	str, err := rg.String()
	if err != nil {
		t.Fatalf("String() failed: %v", err)
	}
	if len(str) != length {
		t.Errorf("Alphabets length mismatch: got %d, want %d", len(str), length)
	}
}

func TestNewChars(t *testing.T) {
	length := 12
	rg, err := NewChars(nil, length)
	if err != nil {
		t.Fatalf("NewChars failed: %v", err)
	}
	str, err := rg.String()
	if err != nil {
		t.Fatalf("String() failed: %v", err)
	}
	if len(str) != length {
		t.Errorf("Chars length mismatch: got %d, want %d", len(str), length)
	}
}

func TestNewAllChars(t *testing.T) {
	length := 16
	rg, err := NewAllChars(nil, length)
	if err != nil {
		t.Fatalf("NewAllChars failed: %v", err)
	}
	str, err := rg.String()
	if err != nil {
		t.Fatalf("String() failed: %v", err)
	}
	if len(str) != length {
		t.Errorf("AllChars length mismatch: got %d, want %d", len(str), length)
	}
}

func TestNewBytes(t *testing.T) {
	length := 20
	rg, err := NewBytes(nil, length)
	if err != nil {
		t.Fatalf("NewBytes failed: %v", err)
	}
	b, err := rg.Bytes()
	if err != nil {
		t.Fatalf("Bytes() failed: %v", err)
	}
	if len(b) != length {
		t.Errorf("Bytes length mismatch: got %d, want %d", len(b), length)
	}
}

func TestNewRandomFromSet(t *testing.T) {
	set := []byte{'a', 'b', 'c'}
	length := 5
	rg, err := NewRandomFromSet(nil, length, set)
	if err != nil {
		t.Fatalf("NewRandomFromSet failed: %v", err)
	}
	b, err := rg.Bytes()
	if err != nil {
		t.Fatalf("Bytes() failed: %v", err)
	}
	if len(b) != length {
		t.Errorf("RandomFromSet length mismatch: got %d, want %d", len(b), length)
	}
	for _, v := range b {
		found := false
		for _, s := range set {
			if v == s {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("RandomFromSet: byte %v not in set %v", v, set)
		}
	}
}
