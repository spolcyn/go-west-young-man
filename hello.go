// all programs start running in package main
package main

import (
    "fmt" //'formatted IO'
)

type Vertex struct {
    X int
    Y int
}

var (
    v1 = Vertex{1,2}
    v2 = Vertex{X: 1} // Y:0 is implicit
    v3 = Vertex{}
    p = &Vertex{1,2} //has type *Vertex
)

func main() {
    fmt.Println(v1, p, v2, v3)
}


