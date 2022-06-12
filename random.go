package mafu

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"time"
)

//RandomInt return a number from min to max - 1
func RandInt(min, max int) (int, error) {
	i, err := Rand0ToInt(max - min)
	if err != nil {
		return min, err
	}
	i += min
	return i, nil
}

//Rand0ToInt return a number from 0 to max - 1, return 0 if max == 0 and return error if max's negative
func Rand0ToInt(max int) (int, error) {
	if max == 0 {
		return 0, nil
	}
	preRand, err := crand.Int(crand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(preRand.Int64()), nil
}

//RandFloat return a number from min to max and return 0 if min > max
func RandFloat(min, max float64) float64 {
	if min > max {
		return min
	}

	rand.Seed(time.Now().UnixNano())
	return (min + rand.Float64()*(max-min))
}
