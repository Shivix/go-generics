package algorithm

import (

)

type Numeric interface {
    int | int8 | int16 | int32 | int64 | float32 | float64
}

type fn[T any] func(T, T) bool

func Max[T Numeric](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func MaxFn[T any](a, b T, comparison fn[T]) T {
    if comparison(a, b) {
        return a
    }
    return b
}

func Min[T Numeric](a, b T) T {
    if a > b {
        return a
    }
    return b
}
