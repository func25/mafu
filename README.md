# mafu

Mafu is a set of simple math utils that I use for my project, which make writing code more convenient (with support of generic 1.18):

- Max, min, random
- Bit processing for enum flag
- Randomly select an element according to the ratio (or the expect rate)
- ...

## Content:

- [Installation](#installation)
- [Samples](#samples)
  - [Find min, max](#find-min-max)
  - [Enum flag](#enum-flag)
  - [Random number](#random-number)
  - [Random with rarities](#random-with-rarities)
  - [Random with expect](#random-with-expect)

## Installation

1. You first need Go installed (version 1.18+ is required):

```sh
$ go get -u github.com/func25/mafu
```

2. Import it in your code:

```go
import "github.com/func25/mafu"
```

## Samples

### Find min, max:

Support ellipsis (...) with all primitive number types: int (int, uint, int32,...), float (float32, float64):

```go
mafu.Min(58, 39, 24, 68, 34)            // 24
mafu.Min(58.5, 39.1, 24.4, 68.67, 34.3) // 24.4
mafu.Max(58.5, 39.1, 24.4, 68.67, 34.3) // 68.67
```

### Enum flag:

this flag is used when you want a variable that can store multiple states

```go
type MultiColor int

const (
	ColorBlack MultiColor = 1 << iota
	ColorRed
	ColorGreen
	ColorBlue
)

func main() {
	mix := ColorBlack | ColorGreen | ColorBlue
	fmt.Println(mix) // 13 = 1101

	mafu.FlagHas(mix, ColorGreen)       // true
	mix = mafu.FlagOff(mix, ColorGreen) // 9 = 1001
	mafu.FlagHas(mix, ColorGreen)       // false
}
```

### Random number

Also support all int types (int, uint, int32,...) and float types (float32, float64):

```go
func main() {
	mafu.Rand0ToInt(100)        // 6
	mafu.RandFloat(0.0, 100.0)  // 38.9706330749286
}
```

### Random with rarities:

You may want to use RandRateUnits or RandRarities (return index) to random element in an array with rarities:

```go
func main() {
	quantity := map[string]int{}

	for i := 0; i < 10000; i++ {
		res, _ := mafu.RandRateUnits([]mafu.RateUnit[string]{
			{Key: "a", Rate: 0.1},
			{Key: "ab", Rate: 0.1},
			{Key: "abc", Rate: 0.1},
			{Key: "abcd", Rate: 0.1},
			{Key: "abcde", Rate: 0.6},
		})
		quantity[res.Key]++
	}

	fmt.Println(quantity)
}
// result: map[a:1048 ab:962 abc:993 abcd:1049 abcde:5948]

func main() {
	quantity := map[int]int{}
	for i := 0; i < 10000; i++ {
		res, _ := mafu.RandRarities([]float64{0.1, 0.1, 0.1, 0.1, 0.6})
		quantity[res]++
	}

	fmt.Println(quantity)
}
// result: map[0:1074 1:993 2:1069 3:1086 4:5778]

```

### Random with expect

If you want more precise result, use `RandExpect` algorithm:

```go
func main() {
	showed := map[string]uint{}
	for i := 0; i < 10000; i++ {
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
// result: map[a:1000 ab:1000 abc:1000 abcd:1000 abcde:6000]
```
