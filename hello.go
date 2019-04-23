// all programs start running in package main
package main

import (
    "fmt" //'formatted IO'
    "math"
)

// arguments: a function which takes 2 float64's and returns a float 64
// return val: a float64
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)
}

func main() {
    hypot := func(x, y float64) float64 {
        return math.Sqrt(x*x + y*y)
    }
    fmt.Println(hypot(5, 12))

    fmt.Println(compute(hypot))
    fmt.Println(compute(math.Pow))
}
