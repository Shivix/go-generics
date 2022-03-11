package optional

import(

)

type Optional[T any] struct {
    value T
    flag bool
}

func New[T any](value T) Optional[T] {
    return Optional[T]{
        value: value,
        flag: true,
    }
}

func NullOpt[T any]() Optional[T] {
    var result T
    return Optional[T] {
        value: result,
        flag: false,
    }
}

func HasValue[T any](opt Optional[T]) bool {
    return opt.flag
}

func ValueOr[T any](opt Optional[T], value T) T {
    if opt.flag {
        return opt.value
    }
    return value
}

func Value[T any](opt Optional[T]) T {
    if opt.flag {
        return opt.value
    }
    panic("Optional contained no data")
}

