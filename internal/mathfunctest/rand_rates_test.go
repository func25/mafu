package mafutest

import (
	"fmt"
	"testing"

	"github.com/func25/mafu"
)

func TestRandRates(t *testing.T) {
	quantity := map[string]int{}
	n := 10000
	for i := 0; i < n; i++ {
		res, err := mafu.RandRateUnits([]mafu.RateUnit[string]{
			{Key: "a", Rate: 0.1},
			{Key: "ab", Rate: 0.1},
			{Key: "abc", Rate: 0.1},
			{Key: "abcd", Rate: 0.1},
			{Key: "abcde", Rate: 0.6},
		})
		if err != nil {
			t.Error(err)
			return
		}
		quantity[res.Key]++
	}

	fmt.Println(quantity)
}

func TestRandRarity(t *testing.T) {
	quantity := map[int]int{}
	n := 10000
	for i := 0; i < n; i++ {
		res, err := mafu.RandRarities([]float64{0.1, 0.1, 0.1, 0.1, 0.6})
		if err != nil {
			t.Error(err)
			return
		}
		quantity[res]++
	}

	fmt.Println(quantity)
}
