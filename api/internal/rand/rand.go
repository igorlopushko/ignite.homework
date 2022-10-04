// Package rand contains utility methods for the random numbers generation.
package rand

import (
	"crypto/rand"
	"math/big"
)

// Gets a random number in rage of 0 to N.
func GetRandomNumber(n int) (int64, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		return 0, err
	}
	return r.Int64(), nil
}
