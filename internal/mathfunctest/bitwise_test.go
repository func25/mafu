package mafutest

import (
	"fmt"
	"testing"

	"github.com/func25/mafu"
)

type abc struct {
	a, b, c mafu.UIntBits
}

func TestBitwise(t *testing.T) {
	turnOnTable := []abc{
		{2, 1, 3},
		{89, 20, 93},
	}
	t.Run("turn on", func(t *testing.T) {
		for _, v := range turnOnTable {
			if v.a.On(v.b) != v.c {
				t.Error(fmt.Sprintf("turn on failed case: %v", v))
			}
		}
	})

	turnOffTable := []abc{
		{2, 1, 2},
		{2, 2, 0},
		{89, 20, 73},
	}
	t.Run("turn off", func(t *testing.T) {
		for _, v := range turnOffTable {
			if v.a.Off(v.b) != v.c {
				t.Error(fmt.Sprintf("turn on failed case: %v", v))
			}
		}
	})

	hasTable := []abc{
		{2, 1, 0},
		{3, 1, 1},
		{3, 2, 1},
		{45, 36, 1},
		{45, 35, 0},
	}
	t.Run("turn off", func(t *testing.T) {
		for _, v := range hasTable {
			res := true
			if v.c == 0 {
				res = false
			}
			if v.a.Has(v.b) != res {
				t.Error(fmt.Sprintf("turn on failed case: %v", v))
			}
		}
	})
}
