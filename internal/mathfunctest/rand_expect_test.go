package mafutest

import (
	"fmt"
	"testing"

	"github.com/func25/mafu"
)

func TestExpect(t *testing.T) {
	n := 10000
	showed := map[string]uint{}
	for i := 0; i < n; i++ {
		res, _ := mafu.RandExpect([]mafu.ExpectUnit[string]{
			{Key: "a", Expect: 0.1},
			{Key: "ab", Expect: 0.1},
			{Key: "abc", Expect: 0.1},
			{Key: "abcd", Expect: 0.1},
			{Key: "abcde", Expect: 0.6},
		}, showed)

		showed[res.Key]++
	}

	fmt.Println(showed)
}
