package util

import "github.com/jigadhirasu/rex"

func TransferToFunc1[A, B any](f rex.Transfer1[A, B]) rex.Func1[A, B] {
	return func(ctx rex.Context, a A) (B, error) {
		return f(a), nil
	}
}

func TransferToFunc2[A, B, C any](f rex.Transfer2[A, B, C]) rex.Func2[A, B, C] {
	return func(ctx rex.Context, a A, b B) (C, error) {
		return f(a, b), nil
	}
}
