package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var a [][]uint8
	a = make([][]uint8, dy)
	var b []uint8
	for i := 0; i < dy; i++ {
		b = make([]uint8, dx)
		a[i] = b
		for j := 0; j < dx; j++ {
			a[i][j] = uint8(i) * uint8(100)
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
