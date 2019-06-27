// all programs start running in package main
package main

// Exercise: Equivalent Binary Trees
// Want to check whether two binary trees store same sequence of values
// Uses tree package, which defines type Tree struct { Left *Tree Value int Right *Tree }
// We'll implement Walk & Same functions

import (
    "golang.org/x/tour/tree"
    "fmt"
)

// `Walk` walks the tree t, sending all values
// from the tree to the channel ch
// Inorder traversal
func Walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }

    Walk(t.Left, ch)
    ch <- t.Value
    Walk(t.Right, ch)
}

// `Same` determines whether the trees t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int, 10)
    ch2 := make(chan int, 10)

    go Walk(t1, ch1)
    go Walk(t2, ch2)

    for {
        i, closed1 := <- ch1
        j, closed2 := <- ch2
        if i != j {
            return false
        }
        if closed1 && closed2 {
            return true
        }
    }
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
}
