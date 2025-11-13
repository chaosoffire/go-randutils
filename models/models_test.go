package models

import (
	"slices"
	"testing"
)

// TestNumset verifies the Numset contains correct ASCII codes for digits
func TestNumset(t *testing.T) {
	if len(Numset) != 10 {
		t.Errorf("Numset length = %d, want 10", len(Numset))
	}
	expected := []int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57} // '0' to '9'
	if !slices.Equal(Numset, expected) {
		t.Errorf("Numset = %v, want %v", Numset, expected)
	}
}

// TestLowerset verifies the Lowerset contains correct ASCII codes for lowercase letters
func TestLowerset(t *testing.T) {
	if len(Lowerset) != 26 {
		t.Errorf("Lowerset length = %d, want 26", len(Lowerset))
	}
	// Check first and last elements
	if Lowerset[0] != 97 { // 'a'
		t.Errorf("Lowerset[0] = %d, want 97", Lowerset[0])
	}
	if Lowerset[25] != 122 { // 'z'
		t.Errorf("Lowerset[25] = %d, want 122", Lowerset[25])
	}
	// Check all are consecutive
	for i := 0; i < len(Lowerset); i++ {
		if Lowerset[i] != 97+i {
			t.Errorf("Lowerset[%d] = %d, want %d", i, Lowerset[i], 97+i)
		}
	}
}

// TestUpperset verifies the Upperset contains correct ASCII codes for uppercase letters
func TestUpperset(t *testing.T) {
	if len(Upperset) != 26 {
		t.Errorf("Upperset length = %d, want 26", len(Upperset))
	}
	// Check first and last elements
	if Upperset[0] != 65 { // 'A'
		t.Errorf("Upperset[0] = %d, want 65", Upperset[0])
	}
	if Upperset[25] != 90 { // 'Z'
		t.Errorf("Upperset[25] = %d, want 90", Upperset[25])
	}
	// Check all are consecutive
	for i := 0; i < len(Upperset); i++ {
		if Upperset[i] != 65+i {
			t.Errorf("Upperset[%d] = %d, want %d", i, Upperset[i], 65+i)
		}
	}
}

// TestSymbolset verifies the Symbolset contains correct ASCII codes for special symbols
func TestSymbolset(t *testing.T) {
	// 33-47 (15 chars) + 58-64 (7 chars) = 22 chars
	expectedLen := 15 + 7
	if len(Symbolset) != expectedLen {
		t.Errorf("Symbolset length = %d, want %d", len(Symbolset), expectedLen)
	}

	// Check first range (33-47)
	for i := 0; i < 15; i++ {
		if Symbolset[i] != 33+i {
			t.Errorf("Symbolset[%d] = %d, want %d", i, Symbolset[i], 33+i)
		}
	}

	// Check second range (58-64)
	for i := 0; i < 7; i++ {
		if Symbolset[15+i] != 58+i {
			t.Errorf("Symbolset[%d] = %d, want %d", 15+i, Symbolset[15+i], 58+i)
		}
	}
}

// TestAlphabetset verifies the Alphabetset contains all letters
func TestAlphabetset(t *testing.T) {
	expectedLen := len(Upperset) + len(Lowerset)
	if len(Alphabetset) != expectedLen {
		t.Errorf("Alphabetset length = %d, want %d", len(Alphabetset), expectedLen)
	}

	// Verify it's the concatenation of Upperset and Lowerset
	expected := slices.Concat(Upperset, Lowerset)
	if !slices.Equal(Alphabetset, expected) {
		t.Errorf("Alphabetset is not the concatenation of Upperset and Lowerset")
	}
}

// TestCharset verifies the Charset contains all alphanumeric characters
func TestCharset(t *testing.T) {
	expectedLen := len(Numset) + len(Upperset) + len(Lowerset)
	if len(Charset) != expectedLen {
		t.Errorf("Charset length = %d, want %d", len(Charset), expectedLen)
	}

	// Verify it's the concatenation of Numset, Upperset, and Lowerset
	expected := slices.Concat(Numset, Upperset, Lowerset)
	if !slices.Equal(Charset, expected) {
		t.Errorf("Charset is not the concatenation of Numset, Upperset, and Lowerset")
	}
}

// TestAllset verifies the Allset contains all characters (digits, letters, and symbols)
func TestAllset(t *testing.T) {
	expectedLen := len(Numset) + len(Upperset) + len(Lowerset) + len(Symbolset)
	if len(Allset) != expectedLen {
		t.Errorf("Allset length = %d, want %d", len(Allset), expectedLen)
	}

	// Verify it's the concatenation of all sets
	expected := slices.Concat(Numset, Upperset, Lowerset, Symbolset)
	if !slices.Equal(Allset, expected) {
		t.Errorf("Allset is not the concatenation of all sets")
	}
}

// TestCharacterSetCompleteness verifies there are no overlaps between character sets
func TestCharacterSetCompleteness(t *testing.T) {
	tests := []struct {
		name   string
		set1   []int
		set2   []int
		hasErr bool
	}{
		{"Numset vs Upperset", Numset, Upperset, false},
		{"Numset vs Lowerset", Numset, Lowerset, false},
		{"Numset vs Symbolset", Numset, Symbolset, false},
		{"Upperset vs Lowerset", Upperset, Lowerset, false},
		{"Upperset vs Symbolset", Upperset, Symbolset, false},
		{"Lowerset vs Symbolset", Lowerset, Symbolset, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, val1 := range tt.set1 {
				for _, val2 := range tt.set2 {
					if val1 == val2 {
						t.Errorf("%s: found overlap at value %d", tt.name, val1)
					}
				}
			}
		})
	}
}

// TestCharacterRanges verifies the character ranges are correct
func TestCharacterRanges(t *testing.T) {
	tests := []struct {
		name  string
		set   []int
		start int
		end   int
	}{
		{"Numset", Numset, 48, 57},
		{"Upperset", Upperset, 65, 90},
		{"Lowerset", Lowerset, 97, 122},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedLen := tt.end - tt.start + 1
			if len(tt.set) != expectedLen {
				t.Errorf("%s length = %d, want %d", tt.name, len(tt.set), expectedLen)
			}
			for i, val := range tt.set {
				expected := tt.start + i
				if val != expected {
					t.Errorf("%s[%d] = %d, want %d", tt.name, i, val, expected)
				}
			}
		})
	}
}

// TestInitialization verifies that all character sets are properly initialized
func TestInitialization(t *testing.T) {
	// This test ensures that the init() function runs correctly by checking
	// that all sets are non-nil and have the correct sizes
	sets := map[string][]int{
		"Numset":      Numset,
		"Lowerset":    Lowerset,
		"Upperset":    Upperset,
		"Symbolset":   Symbolset,
		"Alphabetset": Alphabetset,
		"Charset":     Charset,
		"Allset":      Allset,
	}

	for name, set := range sets {
		if len(set) == 0 {
			t.Errorf("%s not initialized", name)
		}
	}
}
