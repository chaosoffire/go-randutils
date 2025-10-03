package randutils

// precision for big.Float operations
const (
	precision         = 256
	defaultBufferSize = 256
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
