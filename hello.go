// all programs start running in package main
package main

import (
    "fmt" //'formatted IO'
)

// TODO: Implement a fibonnaci function that returns a function (a closure) that returns successive fibonacci numbers

// fibonacci is a function that returns
// a function that returns an int
func fibonacci() func() int {
    current := 0
    previous := 1

    return func() int {
       temp := current
       current += previous
       previous = temp
       return previous
   }

}

func main() {
    f := fibonacci()

    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
