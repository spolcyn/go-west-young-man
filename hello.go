// all programs start running in package main
package main

import (
    "fmt" //'formatted IO'
)

// arguments: none
// return: f: int -> int
func adder() func(int) int {
    sum := 0

    return func(x int) int {
        sum += x
        return sum
    }

}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
}
