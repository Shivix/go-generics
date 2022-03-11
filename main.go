package main

import(
    "fmt"

    "generics/optional"
    "generics/algorithm"
)

func testopt(input bool) optional.Optional[int] {
    if input {
        return optional.New(5)
    }
    return optional.NullOpt[int]()
}

func main() {
    // Optional
    fmt.Printf("%+v\n",testopt(true))
    opt := optional.New(5)
    fmt.Printf("%d\n", optional.ValueOr(opt, 10))

    // Algorithm
    fmt.Printf("%d\n", algorithm.Max(15, 51))
    fmt.Printf("%d\n", algorithm.MaxFn(15, 51, func(a, b int) bool { return a > b }))
}
