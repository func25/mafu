// this file is a portion extracted from https://pkg.go.dev/golang.org/x/exp/constraints if you're interested
package mafu

type Ordered interface {
	Integer | Float | ~string
}

type Number interface {
	Integer | Float
}

type SignedNumber interface {
	Signed | Float
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

type Float interface {
	~float32 | ~float64
}
