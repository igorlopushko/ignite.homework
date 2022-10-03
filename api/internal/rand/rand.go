// Package rand contains utility methods for the random numbers generation
package rand

import (
	"crypto/rand"
	"math/big"
)

func GetRandomNumber(l int) (int64, error) {
	r, err := rand.Int(rand.Reader, big.NewInt(int64(l)))
	if err != nil {
		return 0, err
	}
	return r.Int64(), nil
}
