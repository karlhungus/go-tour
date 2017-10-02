package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	square := make([][]uint8, dy)
	for i, _ := range square {
		square[i] = make([]uint8, dx)
		for j, _ := range square[i] {
			square[i][j] = uint8(i + j/2)
		}
	}
	return square

}

func main() {
	pic.Show(Pic)
}
