// Package models defines the character sets used for random generation.
package models

import "slices"

// ASCII ranges for character categories
const (
	numStart     = 48  // '0'
	numEnd       = 57  // '9'
	lowerStart   = 97  // 'a'
	lowerEnd     = 122 // 'z'
	upperStart   = 65  // 'A'
	upperEnd     = 90  // 'Z'
	symbol1Start = 33  // '!'
	symbol1End   = 47  // '/'
	symbol2Start = 58  // ':'
	symbol2End   = 64  // '@'
)

var (
	// Numset contains ASCII codes for digits 0-9
	Numset = make([]int, numEnd-numStart+1)
	// Lowerset contains ASCII codes for lowercase letters a-z
	Lowerset = make([]int, lowerEnd-lowerStart+1)
	// Upperset contains ASCII codes for uppercase letters A-Z
	Upperset = make([]int, upperEnd-upperStart+1)
	// Symbolset contains ASCII codes for special symbols (!-/ and :-@)
	Symbolset = make([]int, (symbol1End-symbol1Start+1)+(symbol2End-symbol2Start+1))

	// Alphabetset contains ASCII codes for all letters (A-Z, a-z)
	Alphabetset = make([]int, len(Lowerset)+len(Upperset))
	// Charset contains ASCII codes for alphanumeric characters (0-9, A-Z, a-z)
	Charset = make([]int, len(Numset)+len(Lowerset)+len(Upperset))
	// Allset contains ASCII codes for all characters: digits, letters, and symbols
	Allset = make([]int, len(Numset)+len(Lowerset)+len(Upperset)+len(Symbolset))
)

func init() {
	// Initialize digit ASCII codes (48-57)
	for i := range numEnd - numStart + 1 {
		Numset[i] = numStart + i
	}
	// Initialize lowercase letter ASCII codes (97-122)
	for i := range lowerEnd - lowerStart + 1 {
		Lowerset[i] = lowerStart + i
	}
	// Initialize uppercase letter ASCII codes (65-90)
	for i := range upperEnd - upperStart + 1 {
		Upperset[i] = upperStart + i
	}
	// Initialize first symbol range ASCII codes (33-47)
	for i := range symbol1End - symbol1Start + 1 {
		Symbolset[i] = symbol1Start + i
	}
	// Initialize second symbol range ASCII codes (58-64)
	for i := range symbol2End - symbol2Start + 1 {
		Symbolset[(symbol1End-symbol1Start+1)+i] = symbol2Start + i
	}

	// Combine character sets using slices.Concat
	Alphabetset = slices.Concat(Upperset, Lowerset)
	Charset = slices.Concat(Numset, Upperset, Lowerset)
	Allset = slices.Concat(Numset, Upperset, Lowerset, Symbolset)
}
