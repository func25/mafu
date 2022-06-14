package mafu

import (
	"errors"
)

var (
	ErrEmpty  = errors.New("one or more parameters are empty")
	ErrNilKey = errors.New("(one of the parameters)'s key is nil")
)

// -------------- Rarity strategy

//RateUnit expect rate must be in [0, 1]
type RateUnit[T any] struct {
	Key  T       `json:"key"`
	Rate float64 `json:"rate"`
}

//RandRarities return the index of selected element in rateUnits array
func RandRarities(rateUnits []float64) (int, error) {
	if len(rateUnits) <= 0 {
		return 0, errors.New("cannot rate pick the empty array")
	}

	var totalRate float64 = 0
	for _, rate := range rateUnits {
		totalRate += rate
	}

	randNum := RandFloat(0, totalRate)

	var lastID int = 0
	for k, v := range rateUnits {
		var percent float64 = v
		if percent > 0 {
			lastID = k
			if randNum > percent {
				randNum -= percent
			} else {
				break
			}
		}
	}

	return lastID, nil
}

func RandRateUnits[T any](units []RateUnit[T]) (resp RandResult[T], err error) {
	if len(units) <= 0 {
		return resp, errors.New("cannot rate pick the empty array")
	}

	var totalRate float64 = 0
	for _, unit := range units {
		totalRate += unit.Rate
	}

	randNum := RandFloat(0, totalRate)

	last := units[0].Key
	id := 0
	for k := range units {
		var percent float64 = units[k].Rate
		if percent > 0 {
			last = units[k].Key
			id = k
			if randNum > percent {
				randNum -= percent
			} else {
				break
			}
		}
	}

	return RandResult[T]{
		Key: last,
		ID:  id,
	}, nil
}

// ---------------- Expect strategy
type ExpectUnit[T comparable] struct {
	Key    T       `json:"key"`
	Expect float64 `json:"expect"`
}

type RandResult[T any] struct {
	Key T   `json:"key"`
	ID  int `json:"id"`
}

// options: first get first
// options: random selected
// options: update showed
// options: balance showed
func RandExpect[T comparable](units []ExpectUnit[T], showed map[T]int) (res RandResult[T], err error) {
	if len(units) <= 0 {
		return res, ErrEmpty
	}

	totalShowed := 0
	var totalExpect float64 = 0
	for i := range units {
		totalShowed += showed[units[i].Key]
		totalExpect += units[i].Expect
	}

	if totalShowed == 0 {
		id, err := Rand0ToInt(len(units))
		if err != nil {
			return res, err
		}
		return RandResult[T]{
			Key: units[id].Key,
			ID:  id,
		}, nil
	}

	actuals := make([]int, len(units))
	minActual := float64(showed[units[0].Key])/float64(totalShowed)*totalExpect - units[0].Expect
	count := 0
	for k, v := range units {
		actual := float64(showed[v.Key])/float64(totalShowed)*totalExpect - v.Expect
		if actual > minActual {
			continue
		} else if actual < minActual {
			count = 0
			minActual = actual
		}

		actuals[count] = k
		count++
	}

	id, err := Rand0ToInt(count)
	if err != nil {
		return res, err
	}

	return RandResult[T]{
		Key: units[actuals[id]].Key,
		ID:  actuals[id],
	}, err
}
