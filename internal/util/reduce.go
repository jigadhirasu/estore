package util

import (
	"github.com/jigadhirasu/rex"
)

func ReduceSlice[A any](b []A, a A) []A {
	if b == nil {
		b = []A{}
	}

	b = append(b, a)
	return b
}

func ReduceSliceFunc[A, B any](f rex.Transfer1[A, B]) func(b []B, a A) []B {
	return func(b []B, a A) []B {
		if b == nil {
			b = []B{}
		}

		b = append(b, f(a))
		return b
	}
}
