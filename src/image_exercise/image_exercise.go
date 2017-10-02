package main

import (
	"github.com/karlhungus/lib_pxl"
	"image"
	"image/color"
)

type Image struct{}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 40, 40)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x * y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{}
	//pic.ShowImage(m)
	lib_pxl.Init()
	lib_pxl.Display(m)
	defer lib_pxl.Close()
}
