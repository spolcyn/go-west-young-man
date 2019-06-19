// all programs start running in package main
package main

import (
    "golang.org/x/tour/pic"
    "image"
    "image/color"
)

type Image struct{}

func (m Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
    return image.Rect(0,0, 512, 512)
}

func (m Image) At(x, y int) color.Color {
    return color.RGBA{uint8(x % 255), uint8(y % 255), 255, 255}
}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
