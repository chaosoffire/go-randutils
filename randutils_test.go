package randutils

import (
	"regexp"
	"strings"
	"testing"

	"github.com/chaosoffire/go-randutils/models"
)

// TestInt tests the Int function
func TestInt(t *testing.T) {
	tests := []struct {
		name    string
		max     int
		wantErr bool
	}{
		{"valid positive max", 100, false},
		{"max of 1", 1, false},
		{"invalid max zero", 0, true},
		{"invalid max negative", -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Int(tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && (result < 0 || result >= tt.max) {
				t.Errorf("Int() result = %d, want result in range [0, %d)", result, tt.max)
			}
		})
	}
}

// TestInt_Randomness tests that Int returns different values on multiple calls
func TestInt_Randomness(t *testing.T) {
	const iterations = 100
	max := 1000
	results := make(map[int]bool)

	for i := 0; i < iterations; i++ {
		result, err := Int(max)
		if err != nil {
			t.Fatalf("Int() failed: %v", err)
		}
		results[result] = true
	}

	// With high probability, we should get at least some variety
	if len(results) < 50 {
		t.Errorf("Int() produced too few unique values: %d out of %d iterations", len(results), iterations)
	}
}

// TestIntRange tests the IntRange function
func TestIntRange(t *testing.T) {
	tests := []struct {
		name    string
		min     int
		max     int
		wantErr bool
	}{
		{"valid range", 10, 20, false},
		{"valid range single value", 5, 6, false},
		{"invalid range equal", 10, 10, true},
		{"invalid range reversed", 20, 10, true},
		{"negative range valid", -10, -5, false},
		{"negative to positive", -5, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := IntRange(tt.min, tt.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntRange(%d, %d) error = %v, wantErr %v", tt.min, tt.max, err, tt.wantErr)
			}
			if !tt.wantErr && (result < tt.min || result >= tt.max) {
				t.Errorf("IntRange(%d, %d) result = %d, want result in range [%d, %d)", tt.min, tt.max, result, tt.min, tt.max)
			}
		})
	}
}

// TestRandom tests the Random function
func TestRandom(t *testing.T) {
	tests := []struct {
		name     string
		length   int
		charset  []int
		wantErr  bool
		checkErr bool
	}{
		{"valid random", 10, models.Numset, false, false},
		{"valid random alphabet", 20, models.Alphabetset, false, false},
		{"single element", 1, []int{65}, false, false},
		{"invalid length zero", 0, models.Numset, true, false},
		{"invalid length negative", -5, models.Numset, true, false},
		{"empty charset", 10, []int{}, true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Random(tt.length, tt.charset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Random() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && len(result) != tt.length {
				t.Errorf("Random() length = %d, want %d", len(result), tt.length)
			}
			if !tt.wantErr {
				// Verify all values are in the charset
				for _, val := range result {
					found := false
					for _, ch := range tt.charset {
						if val == ch {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("Random() returned value %d not in charset", val)
					}
				}
			}
		})
	}
}

// TestStrings tests the Strings function
func TestStrings(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"valid 10 chars", 10, false},
		{"valid 100 chars", 100, false},
		{"valid single char", 1, false},
		{"invalid length zero", 0, true},
		{"invalid length negative", -5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Strings(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Strings() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if len(result) != tt.length {
					t.Errorf("Strings() length = %d, want %d", len(result), tt.length)
				}
				// Verify all characters are alphanumeric (digits, uppercase, lowercase)
				for _, ch := range result {
					if !((ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')) {
						t.Errorf("Strings() returned non-alphanumeric character: %c", ch)
					}
				}
			}
		})
	}
}

// TestByte tests the Byte function
func TestByte(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"valid 16 bytes", 16, false},
		{"valid 256 bytes", 256, false},
		{"valid single byte", 1, false},
		{"invalid length zero", 0, true},
		{"invalid length negative", -10, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Byte(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Byte() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && len(result) != tt.length {
				t.Errorf("Byte() length = %d, want %d", len(result), tt.length)
			}
		})
	}
}

// TestByte_Randomness tests that Byte produces different values on multiple calls
func TestByte_Randomness(t *testing.T) {
	const iterations = 10
	length := 16
	results := make(map[string]bool)

	for i := 0; i < iterations; i++ {
		result, err := Byte(length)
		if err != nil {
			t.Fatalf("Byte() failed: %v", err)
		}
		results[string(result)] = true
	}

	if len(results) != iterations {
		t.Errorf("Byte() produced duplicate values, expected %d unique, got %d", iterations, len(results))
	}
}

// TestBase64 tests the Base64 function
func TestBase64(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"valid 16 bytes", 16, false},
		{"valid 32 bytes", 32, false},
		{"valid single byte", 1, false},
		{"invalid length zero", 0, true},
		{"invalid length negative", -5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Base64(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base64() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				// Verify it's valid base64
				validBase64 := regexp.MustCompile(`^[A-Za-z0-9+/]*={0,2}$`)
				if !validBase64.MatchString(result) {
					t.Errorf("Base64() returned invalid base64: %s", result)
				}
				// Check approximate length (base64 adds padding)
				expectedLen := (tt.length*4 + 2) / 3
				if len(result) < expectedLen || len(result) > expectedLen+2 {
					t.Logf("Base64() length = %d, expected approximately %d", len(result), expectedLen)
				}
			}
		})
	}
}

// TestHex tests the Hex function
func TestHex(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"valid 16 bytes", 16, false},
		{"valid 32 bytes", 32, false},
		{"valid single byte", 1, false},
		{"invalid length zero", 0, true},
		{"invalid length negative", -5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Hex(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hex() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				// Verify it's valid hex
				if !regexp.MustCompile(`^[0-9a-f]*$`).MatchString(result) {
					t.Errorf("Hex() returned invalid hex: %s", result)
				}
				// Each byte produces 2 hex characters
				if len(result) != tt.length*2 {
					t.Errorf("Hex() length = %d, want %d", len(result), tt.length*2)
				}
			}
		})
	}
}

// TestUUID tests the UUID function
func TestUUID(t *testing.T) {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

	for i := 0; i < 10; i++ {
		t.Run("uuid_"+string(rune(i)), func(t *testing.T) {
			result, err := UUID()
			if err != nil {
				t.Errorf("UUID() error = %v", err)
			}
			if !uuidRegex.MatchString(strings.ToLower(result)) {
				t.Errorf("UUID() returned invalid UUID format: %s", result)
			}
		})
	}
}

// TestUUID_Uniqueness tests that UUID generates unique values
func TestUUID_Uniqueness(t *testing.T) {
	const iterations = 100
	uuids := make(map[string]bool)

	for i := 0; i < iterations; i++ {
		uuid, err := UUID()
		if err != nil {
			t.Fatalf("UUID() failed: %v", err)
		}
		if uuids[uuid] {
			t.Errorf("UUID() generated duplicate: %s", uuid)
		}
		uuids[uuid] = true
	}

	if len(uuids) != iterations {
		t.Errorf("UUID() produced duplicate values, expected %d unique, got %d", iterations, len(uuids))
	}
}

// TestAllChars tests the AllChars function
func TestAllChars(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantErr bool
	}{
		{"valid 20 chars", 20, false},
		{"valid 100 chars", 100, false},
		{"valid single char", 1, false},
		{"invalid length zero", 0, true},
		{"invalid length negative", -5, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := AllChars(tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllChars() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if len(result) != tt.length {
					t.Errorf("AllChars() length = %d, want %d", len(result), tt.length)
				}
				// Verify all characters are in the expected set
				for _, ch := range result {
					charCode := int(ch)
					found := false
					for _, val := range models.Allset {
						if charCode == val {
							found = true
							break
						}
					}
					if !found {
						t.Errorf("AllChars() returned unexpected character: %c (code %d)", ch, charCode)
					}
				}
			}
		})
	}
}

// TestToASCII tests the toASCII function
func TestToASCII(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{"simple digits", []int{48, 49, 50}, "012"},
		{"uppercase letters", []int{65, 66, 67}, "ABC"},
		{"lowercase letters", []int{97, 98, 99}, "abc"},
		{"empty slice", []int{}, ""},
		{"mixed ASCII", []int{72, 101, 108, 108, 111}, "Hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toASCII(tt.input)
			if result != tt.expected {
				t.Errorf("toASCII() = %s, want %s", result, tt.expected)
			}
		})
	}
}
