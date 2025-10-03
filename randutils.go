package randutils

import (
	"fmt"
)

func init() {
	var err error
	Reader, err = NewRandBufferReader()
	if err != nil {
		panic(fmt.Sprintf("randutils: failed to initialize cryptographic random source: %v", err))
	}
}
