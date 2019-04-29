// all programs start running in package main
package main

import (
    "fmt" //'formatted IO'
)

type I interface {
    M()
}

type T struct {
    S string
}

func (t T) M() {
    fmt.Println(t.S)
}

func main() {
    var i I = T{"hello"}
    i.M()
}
