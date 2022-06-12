package mafutest

import (
	"fmt"
	"testing"

	"github.com/func25/mafu"
)

func TestRandRates(t *testing.T) {
	quantity := map[string]int{}
	n := 1000
	for i := 0; i < n; i++ {
		res, err := mafu.RandRateUnits([]mafu.RateUnit[string]{
			{
				Key:  "a",
				Rate: 0.1,
			},
			{
				Key:  "ab",
				Rate: 0.1,
			},
			{
				Key:  "abc",
				Rate: 0.1,
			},
			{
				Key:  "abcd",
				Rate: 0.1,
			},
			{
				Key:  "abcde",
				Rate: 0.6,
			},
		})
		if err != nil {
			t.Error(err)
			return
		}
		quantity[res]++
	}

	fmt.Println(quantity)
}

type abc struct {
	a, b, c mafu.UIntBits
}

func TestAbc(t *testing.T) {
	// var w io.Writer
	// w = os.Stdout
	// w = new(bytes.Buffer)
	// w = nil
	fmt.Println(string([]byte("")))
}

func TestInterface(t *testing.T) {
	var x interface{}
	var y *int = nil
	x = y

	if x != nil {
		fmt.Println("x != nil")
	} else {
		fmt.Println("x == nil")
	}
}

func TestPointer(t *testing.T) {
	var x *int = new(int)
	y := *x
	fmt.Println(x, &y)
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
