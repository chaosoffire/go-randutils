package randutils

import (
	"crypto/rand"
	"errors"
	"io"
	"sync"
)

type RandBufferReader struct {
	buffer []byte
	index  int
	mu     sync.Mutex
}

func NewRandBufferReader() (*RandBufferReader, error) {
	return NewRandBufferReaderWithSize(defaultBufferSize)
}

func NewRandBufferReaderWithSize(size int) (*RandBufferReader, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}
	result := &RandBufferReader{
		buffer: make([]byte, size),
		index:  size,
	}
	result.mu.Lock()
	defer result.mu.Unlock()
	if err := result.fill(); err != nil {
		return nil, err
	}
	return result, nil
}

// fill is an internal, non-thread-safe function.
// It assumes the caller holds the lock.
func (r *RandBufferReader) fill() error {
	if _, err := io.ReadFull(rand.Reader, r.buffer); err != nil {
		return err
	}
	r.index = 0
	return nil
}

func (r *RandBufferReader) Byte() (byte, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.byteUnsafe()
}

func (r *RandBufferReader) byteUnsafe() (byte, error) {
	if r.index >= len(r.buffer) {
		if err := r.fill(); err != nil {
			return 0, err
		}
	}
	b := r.buffer[r.index]
	r.index++
	return b, nil
}

func (r *RandBufferReader) Bytes(n int) ([]byte, error) {
	if n <= 0 {
		return nil, errors.New("invalid byte slice length")
	}
	result := make([]byte, n)
	byteReads, err := io.ReadFull(r, result)
	if err != nil {
		return nil, err
	}
	return result[:byteReads], nil
}

func (r *RandBufferReader) Read(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, nil
	}
	bytesCopied := 0
	r.mu.Lock()
	defer r.mu.Unlock()
	for bytesCopied < len(b) {
		if r.index >= len(r.buffer) {
			if err := r.fill(); err != nil {
				return bytesCopied, err
			}
		}
		byteAvailable := len(r.buffer) - r.index
		needed := len(b) - bytesCopied
		toCopy := min(needed, byteAvailable)
		copy(b[bytesCopied:], r.buffer[r.index:r.index+toCopy])

		r.index += toCopy
		bytesCopied += toCopy
	}
	return bytesCopied, nil
}

func (r *RandBufferReader) ReadRange(b []byte, rng [2]byte) (int, error) {
	length := len(b)
	if length == 0 {
		return 0, nil
	}
	if rng[0] > rng[1] {
		return 0, errors.New("invalid range")
	}
	bytesRead := 0
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range b {
		for {
			tempByte, err := r.byteUnsafe()
			if err != nil {
				return bytesRead, err
			}
			if tempByte >= rng[0] && tempByte <= rng[1] {
				b[i] = tempByte
				bytesRead++
				break
			}
		}
	}
	return bytesRead, nil
}
