package main

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/tour/pic"
)

type Image struct{}

// ColorModel returns the color model of the image.
func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the bounds of the image.
func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

// At returns the color at a specific pixel (x, y).
func (img Image) At(x, y int) color.Color {
	// Create a unique pattern using sine and cosine functions
	r := uint8((128 + 127*math.Sin(float64(x)*0.1)) / 2)
	g := uint8((128 + 127*math.Cos(float64(y)*0.1)) / 2)
	b := uint8((r + g) / 2)
	return color.RGBA{r, g, b, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
