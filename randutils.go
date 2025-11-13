// Package randutils provides cryptographically secure random generation utilities.
// It offers functions for generating random integers, strings, UUIDs, and encoded random data.
package randutils

import (
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"

	"github.com/chaosoffire/go-randutils/models"
)

// Int returns a random integer in the range [0, max) (min inclusive, max exclusive).
// It uses cryptographic randomness and returns an error if max <= 0 or on crypto/rand failure.
func Int(max int) (int, error) {
	if max <= 0 {
		return 0, fmt.Errorf("invalid max: %d", max)
	}
	result, err := crand.Int(crand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("crypto/rand error: %w", err)
	}
	return int(result.Int64()), nil
}

// IntRange returns a random integer in the range [min, max) (min inclusive, max exclusive).
func IntRange(min, max int) (int, error) {
	if min >= max {
		return 0, fmt.Errorf("invalid range: min %d >= max %d", min, max)
	}
	n := max - min
	result, err := Int(n)
	if err != nil {
		return 0, err
	}
	return result + min, nil
}

// Random generates a random sequence of integers by selecting from the provided charset.
// It uses IntRange to properly sample from the charset, ensuring all elements have equal probability.
func Random(length int, charset []int) ([]int, error) {
	if length <= 0 {
		return nil, fmt.Errorf("invalid length: %d", length)
	}
	lengthSet := len(charset)
	if lengthSet == 0 {
		return nil, fmt.Errorf("charset is empty")
	}
	b := make([]int, 0, length)
	for range length {
		idx, err := IntRange(0, lengthSet)
		if err != nil {
			return nil, err
		}
		b = append(b, charset[idx])
	}
	return b, nil
}

// Strings generates a random string of specified length using alphanumeric characters (A-Z, a-z, 0-9).
// Returns an error if length <= 0 or if random generation fails.
func Strings(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("invalid length: %d", length)
	}
	randomInts, err := Random(length, models.Charset)
	if err != nil {
		return "", err
	}
	return toASCII(randomInts), nil
}

// Byte generates a random byte slice of specified length using cryptographic randomness.
// Returns an error if length <= 0 or if random generation fails.
func Byte(length int) ([]byte, error) {
	if length <= 0 {
		return nil, fmt.Errorf("invalid length: %d", length)
	}
	result := make([]byte, length)
	_, err := io.ReadFull(crand.Reader, result)
	if err != nil {
		return nil, fmt.Errorf("failed to read random bytes: %w", err)
	}
	return result, nil
}

// Base64 generates a random base64-encoded string from random bytes.
// The length parameter specifies the number of random bytes to generate (not the output string length).
// The output string will be approximately 4/3 * length characters due to base64 encoding.
func Base64(length int) (string, error) {
	result, err := Byte(length)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), nil
}

// Hex generates a random hexadecimal string from random bytes.
// The length parameter specifies the number of random bytes to generate.
// The output string will be 2 * length characters (each byte produces 2 hex digits).
func Hex(length int) (string, error) {
	result, err := Byte(length)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", result), nil
}

// UUID generates a random RFC 4122 version 4 UUID.
func UUID() (string, error) {
	b := make([]byte, 16)
	_, err := io.ReadFull(crand.Reader, b)
	if err != nil {
		return "", fmt.Errorf("failed to read random bytes: %w", err)
	}
	// Set the version to 4
	b[6] = (b[6] & 0x0f) | 0x40
	// Set the variant to RFC 4122
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		b[0:4],
		b[4:6],
		b[6:8],
		b[8:10],
		b[10:]), nil
}

// AllChars generates a random string of specified length using:
// - uppercase letters (A-Z)
// - lowercase letters (a-z)
// - digits (0-9)
// - special symbols (!@#$%^&*()_+-=[]{}|;:,.<>?)
// Returns an error if length <= 0 or if random generation fails.
func AllChars(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("invalid length: %d", length)
	}
	randomInts, err := Random(length, models.Allset)
	if err != nil {
		return "", err
	}
	return toASCII(randomInts), nil
}

// toASCII converts a slice of ASCII integer values to a string.
func toASCII(ints []int) string {
	b := make([]byte, len(ints))
	for i, v := range ints {
		b[i] = byte(v)
	}
	return string(b)
}
