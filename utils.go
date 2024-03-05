package enthistory

import (
	"fmt"
)

func deref[T any](t *T) T {
	var zero T
	if t == nil {
		return zero
	}
	return *t
}

func typedSliceToType[T any, U any](records []T) ([]U, error) {
	res := make([]U, len(records))
	for i, record := range records {
		a := any(record)
		if b, ok := a.(U); ok {
			res[i] = b
		} else {
			return nil, fmt.Errorf("failed to convert %T to %T", a, b)
		}
	}
	return res, nil
}

func reduce[T any, R any](collection []T, accumulator func(agg R, item T) R, initial R) R {
	for _, item := range collection {
		initial = accumulator(initial, item)
	}

	return initial
}
