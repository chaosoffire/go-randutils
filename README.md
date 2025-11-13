# go-randutils

A Go package providing cryptographically secure random generation utilities. Generate random integers, strings, UUIDs, and encoded random data with ease.

## Features

- **Cryptographically Secure**: Uses `crypto/rand` for all random generation
- **Random Integers**: Generate random integers within specified ranges
- **Random Strings**: Generate random alphabetic strings or strings with mixed character sets
- **Random Bytes**: Generate random byte slices
- **Encoded Output**: Support for Base64 and Hexadecimal encoding
- **UUID Generation**: Generate RFC 4122 version 4 UUIDs
- **Flexible Character Sets**: Pre-defined character sets for digits, letters, symbols, and combinations

## Installation

```bash
go get github.com/chaosoffire/go-randutils
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	"github.com/chaosoffire/go-randutils"
)

func main() {
	// Generate a random integer between 0 and 100
	num, err := randutils.Int(100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random int:", num)

	// Generate a random integer between 10 and 20
	rangeNum, err := randutils.IntRange(10, 20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random range:", rangeNum)

	// Generate a random alphabetic string
	str, err := randutils.Strings(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random string:", str)

	// Generate a random UUID
	uuid, err := randutils.UUID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random UUID:", uuid)

	// Generate a Base64-encoded random string
	b64, err := randutils.Base64(16)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random Base64:", b64)

	// Generate a hexadecimal string
	hex, err := randutils.Hex(16)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random Hex:", hex)
}
```

## API Reference

### Integer Functions

#### `Int(max int) (int, error)`
Returns a random integer in the range `[0, max)` (min inclusive, max exclusive).

- **Parameters**: `max` - Upper bound (exclusive)
- **Returns**: Random integer or error if `max <= 0`

Example:
```go
num, err := randutils.Int(100)  // 0-99
```

#### `IntRange(min, max int) (int, error)`
Returns a random integer in the range `[min, max)` (min inclusive, max exclusive).

- **Parameters**: `min`, `max` - Range bounds
- **Returns**: Random integer or error if `min >= max`

Example:
```go
num, err := randutils.IntRange(10, 20)  // 10-19
```

### String Functions

#### `Strings(length int) (string, error)`
Generates a random string of specified length using alphanumeric characters (A-Z, a-z, 0-9).

- **Parameters**: `length` - Number of characters
- **Returns**: Random alphanumeric string or error if `length <= 0`

Example:
```go
str, err := randutils.Strings(10)  // "Ab3Cd9EfG2"
```

#### `AllChars(length int) (string, error)`
Generates a random string using all character types:
- Uppercase letters (A-Z)
- Lowercase letters (a-z)
- Digits (0-9)
- Special symbols (!@#$%^&*()_+-=[]{}|;:,.<>?)

- **Parameters**: `length` - Number of characters
- **Returns**: Random mixed-character string or error if `length <= 0`

Example:
```go
str, err := randutils.AllChars(20)  // "aB3!cD9@eF2#gH5$iJ8%"
```

### Byte and Encoding Functions

#### `Byte(length int) ([]byte, error)`
Generates a random byte slice of specified length.

- **Parameters**: `length` - Number of bytes
- **Returns**: Random byte slice or error if `length <= 0`

Example:
```go
bytes, err := randutils.Byte(16)
```

#### `Base64(length int) (string, error)`
Generates a random Base64-encoded string from random bytes.

- **Parameters**: `length` - Number of random bytes to generate (not output string length)
- **Returns**: Base64-encoded string or error if `length <= 0`
- **Note**: Output string will be approximately 4/3 × length characters due to Base64 encoding

Example:
```go
b64, err := randutils.Base64(16)  // "a1b2c3d4e5f6g7h8=="
```

#### `Hex(length int) (string, error)`
Generates a random hexadecimal string from random bytes.

- **Parameters**: `length` - Number of random bytes to generate
- **Returns**: Hexadecimal string or error if `length <= 0`
- **Note**: Output string will be 2 × length characters (each byte produces 2 hex digits)

Example:
```go
hex, err := randutils.Hex(16)  // "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6"
```

### UUID Function

#### `UUID() (string, error)`
Generates a random RFC 4122 version 4 UUID.

- **Returns**: UUID string in format `xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx` or error
- **Format**: Standard UUID format with dashes

Example:
```go
uuid, err := randutils.UUID()  // "550e8400-e29b-41d4-a716-446655440000"
```

### Utility Functions

#### `Random(length int, charset []int) ([]int, error)`
Generates a random sequence of integers by selecting from the provided charset.

- **Parameters**: 
  - `length` - Number of integers to generate
  - `charset` - Slice of available integers to select from
- **Returns**: Random integer slice or error if `length <= 0` or charset is empty

Example:
```go
nums, err := randutils.Random(5, []int{1, 2, 3, 4, 5})
```

## Character Sets (models package)

The `models` package provides pre-defined character sets:

- **Numset**: Digits 0-9 (ASCII 48-57)
- **Lowerset**: Lowercase letters a-z (ASCII 97-122)
- **Upperset**: Uppercase letters A-Z (ASCII 65-90)
- **Symbolset**: Special symbols !-/ and :-@ (ASCII 33-47, 58-64)
- **Alphabetset**: All letters A-Z, a-z
- **Charset**: Alphanumeric characters 0-9, A-Z, a-z
- **Allset**: All characters: digits, letters, and symbols

Example:
```go
import "github.com/chaosoffire/go-randutils/models"

// Use pre-defined character sets
random, err := randutils.Random(10, models.Numset)      // Only digits
random, err := randutils.Random(10, models.Alphabetset) // Only letters
random, err := randutils.Random(10, models.Allset)      // All characters
```

## Error Handling

All functions return an error as the second return value. Common errors include:

- **Invalid length**: `length <= 0` for length-based functions
- **Invalid range**: `min >= max` for range functions
- **Invalid max**: `max <= 0` for Int function
- **Empty charset**: Charset is empty for Random function
- **Crypto/rand failure**: Underlying cryptographic random source failure

Example:
```go
str, err := randutils.Strings(-5)
if err != nil {
	log.Printf("Error: %v", err)  // "invalid length: -5"
}
```

## Security Considerations

- **Cryptographically Secure**: All random generation uses `crypto/rand`, suitable for security-sensitive operations like token generation, session IDs, and cryptographic keys
- **No Weak Randomness**: Never uses `math/rand`, which is not suitable for security purposes
- **Proper Error Handling**: Always check and handle errors, as crypto/rand operations can fail

## Examples

### Generate API Token

```go
token, err := randutils.Base64(32)  // 256-bit token encoded in Base64
if err != nil {
	log.Fatal(err)
}
fmt.Println("API Token:", token)
```

### Generate Session ID

```go
sessionID, err := randutils.Hex(16)  // 128-bit session ID in hexadecimal
if err != nil {
	log.Fatal(err)
}
fmt.Println("Session ID:", sessionID)
```

### Generate Random Password

```go
password, err := randutils.AllChars(16)  // 16-character password with mixed character types
if err != nil {
	log.Fatal(err)
}
fmt.Println("Password:", password)
```

### Generate Multiple UUIDs

```go
for i := 0; i < 5; i++ {
	uuid, err := randutils.UUID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("UUID:", uuid)
}
```

## Testing

The project includes comprehensive test coverage for all functions:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

Tests cover:
- Valid and invalid inputs
- Error conditions
- Randomness verification
- Format validation
- Character set correctness

## Requirements

- Go 1.22.0 or later

## License

See the LICENSE file for license information.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues.

## Support

For issues, questions, or suggestions, please open an issue on GitHub.
