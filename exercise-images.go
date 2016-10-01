package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// https://go-tour-jp.appspot.com/methods/25

type Image struct {
	w, h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (ima Image) At(x, y int) color.Color {
	img_func := func(x, y int) uint8 {
		// return uint8((x + y) / 2)
		return uint8(x ^ y)
	}
	v := img_func(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
