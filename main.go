package main

import (
	"fmt"
	"github.com/hatobus/resistorgen/resistor"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {

	colors, _ := resistor.GenerateColor("4k7")

	width := 200
	height := 100

	resistorimage := image.NewRGBA(image.Rect(0, 0, width, height))
	fillcolor := color.RGBA{255, 192, 203, 255}

	draw.Draw(resistorimage, resistorimage.Bounds(), &image.Uniform{fillcolor}, image.ZP, draw.Src)

	firstband := image.Rect(24, 0, 44, height)
	draw.Draw(resistorimage, firstband, &image.Uniform{colors.FirstBand}, image.ZP, draw.Src)

	secondband := image.Rect(68, 0, 88, height)
	draw.Draw(resistorimage, secondband, &image.Uniform{colors.SecondBand}, image.ZP, draw.Over)

	thirdband := image.Rect(112, 0, 132, height)
	draw.Draw(resistorimage, thirdband, &image.Uniform{colors.ThirdBand}, image.ZP, draw.Over)

	fourthband := image.Rect(156, 0, 176, height)
	draw.Draw(resistorimage, fourthband, &image.Uniform{colors.Tolerance}, image.ZP, draw.Over)

	f, _ := os.Create(fmt.Sprintf("resistor.png"))
	png.Encode(f, resistorimage)
}
