package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    var pic = make([][]uint8, dy)

    for i := 0; i < dy; i++ {
        pic[i] = make([]uint8, dx)

        for j := 0; j < dx; j++ {
            pic[i][j] = uint8(255 - j)
        }
    }

    return pic
}

func main() {
    pic.Show(Pic)
}
