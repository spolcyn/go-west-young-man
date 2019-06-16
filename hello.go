// all programs start running in package main
package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }

    z := x/2
    diff := .000000000000001
    for math.Abs(z*z - x) > diff {
        z -= (z*z - x) / (2*z)
    }

    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
    fmt.Println(ErrNegativeSqrt(-2).Error())
}
