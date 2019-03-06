// all programs start running in package main
package main

import (
    "fmt" //'formatted IO'
    "math"
)

func Sqrt(x float64) float64 {
    
    z:= x/2
    diff := .000000000000001
    for math.Abs(z*z - x) > diff  {
        z -= (z*z - x) / (2*z)
    }

    return z
}

func main() {
    fmt.Println(math.Sqrt(2))
    fmt.Println(Sqrt(2))
}
