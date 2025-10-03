package randutils

import (
	"testing"
)

func TestNewRandBufferReader_DefaultSize(t *testing.T) {
	r, err := NewRandBufferReader()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r == nil {
		t.Fatal("expected non-nil RandBufferReader")
	}
	if len(r.buffer) != defaultBufferSize {
		t.Errorf("expected buffer size %d, got %d", defaultBufferSize, len(r.buffer))
	}
}

func TestNewRandBufferReaderWithSize_InvalidSize(t *testing.T) {
	_, err := NewRandBufferReaderWithSize(0)
	if err == nil {
		t.Fatal("expected error for size 0")
	}
	_, err = NewRandBufferReaderWithSize(-10)
	if err == nil {
		t.Fatal("expected error for negative size")
	}
}

func TestRandBufferReader_Byte(t *testing.T) {
	r, err := NewRandBufferReaderWithSize(8)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for i := 0; i < 20; i++ {
		b, err := r.Byte()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// Just check that we get a byte
		_ = b
	}
}

func TestRandBufferReader_Bytes(t *testing.T) {
	r, err := NewRandBufferReader()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	b, err := r.Bytes(32)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(b) != 32 {
		t.Errorf("expected 32 bytes, got %d", len(b))
	}
	_, err = r.Bytes(0)
	if err == nil {
		t.Error("expected error for n=0")
	}
}

func TestRandBufferReader_Read(t *testing.T) {
	r, err := NewRandBufferReaderWithSize(4)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	buf := make([]byte, 10)
	n, err := r.Read(buf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 10 {
		t.Errorf("expected to read 10 bytes, got %d", n)
	}
}

func TestRandBufferReader_Read_ZeroLen(t *testing.T) {
	r, err := NewRandBufferReader()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	n, err := r.Read([]byte{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 0 {
		t.Errorf("expected 0 bytes read, got %d", n)
	}
}

func TestRandBufferReader_ReadRange(t *testing.T) {
	r, err := NewRandBufferReader()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	buf := make([]byte, 16)
	n, err := r.ReadRange(buf, [2]byte{10, 20})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 16 {
		t.Errorf("expected to read 16 bytes, got %d", n)
	}
	for _, b := range buf {
		if b < 10 || b > 20 {
			t.Errorf("byte %d out of range [10,20]", b)
		}
	}
}

func TestRandBufferReader_ReadRange_InvalidRange(t *testing.T) {
	r, err := NewRandBufferReader()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = r.ReadRange(make([]byte, 5), [2]byte{20, 10})
	if err == nil {
		t.Error("expected error for invalid range")
	}
}

func TestRandBufferReader_ReadRange_ZeroLen(t *testing.T) {
	r, err := NewRandBufferReader()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	n, err := r.ReadRange([]byte{}, [2]byte{0, 255})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if n != 0 {
		t.Errorf("expected 0 bytes read, got %d", n)
	}
}

// Helper to test min function if it's not exported
func TestMinFunction(t *testing.T) {
	if min(1, 2) != 1 {
		t.Error("min(1,2) should be 1")
	}
	if min(2, 1) != 1 {
		t.Error("min(2,1) should be 1")
	}
	if min(2, 2) != 2 {
		t.Error("min(2,2) should be 2")
	}
}

// If min is not defined, define it here for test build.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
