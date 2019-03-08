// all programs start running in package main
package main

import (
    "fmt" //'formatted IO'
)

func main() {
    defer fmt.Println("world")

    fmt.Println("hello")
}
