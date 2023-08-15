package util

import (
	"crypto/rand"
	"math/big"
	fallback "math/rand"
	"time"
)

func Random(a, b int64) int64 {
	v, err := rand.Int(rand.Reader, big.NewInt(b+1))
	if err != nil {
		fallback.Seed(time.Now().UnixNano())
		return int64(fallback.Intn(int(b-a+1)) + int(a))
	}

	return v.Int64()
}
