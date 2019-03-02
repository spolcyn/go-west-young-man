// all programs start running in package main
package main

import (
    "fmt" //'formatted IO'
    "math"
)

func main() {
    var x, y int = 3,4
    var f float64 = math.Sqrt(float64(x*x + y*y))
    var z uint = uint(f)
    i := 42
    j := float64(i)
    k := uint(j)
    fmt.Println(x,y,z)
}
