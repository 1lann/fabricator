package fabricator

import (
	"crypto/rand"
	"math/big"
)

const resolution = 2000000

func randomRange(lower float64, upper float64) float64 {
	r := upper - lower
	num, err := rand.Int(rand.Reader, big.NewInt(resolution))
	if err != nil {
		panic(err)
	}

	return lower + (float64(num.Int64())/resolution)*r
}
