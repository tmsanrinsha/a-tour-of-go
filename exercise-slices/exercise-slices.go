package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)

	lenY := make([]int, dy)
	lenX := make([]int, dx)
	for y := range lenY {
		imgX := make([]uint8, dx)
		for x := range lenX {
			imgX[x] = uint8((x + y) / 2)
		}
		img[y] = imgX
	}
	return img
}

func main() {
	pic.Show(Pic)
}
