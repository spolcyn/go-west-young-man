// all programs start running in package main
package main

import (
    "fmt" //'formatted IO'
)

func main() {
    i, j := 42, 2701

    p := &i
    fmt.Println(*p)
    *p = 21
    fmt.Println(i)

    p = &j
    *p = *p / 37
    fmt.Println(j)
}

