package mafutest

import (
	"fmt"
	"testing"

	"github.com/func25/mafu"
)

func TestMinMaxIE(t *testing.T) {
	mafu.Min(58, 39, 24, 68, 34)            // 24
	mafu.Min(58.5, 39.1, 24.4, 68.67, 34.3) // 24.4
	mafu.Max(58.5, 39.1, 24.4, 68.67, 34.3) // 68.67
}

func TestRand(t *testing.T) {
	fmt.Println(mafu.Rand0ToInt(100)) // 24
	fmt.Println(mafu.RandFloat(0.0, 100.0))
}

//49 <nil>
// 56.45570425829678
