package resistor

import (
	"image/color"
)

type Colorbar struct {
	FirstBand  *color.RGBA
	SecondBand *color.RGBA
	ThirdBand  *color.RGBA
	Tolerance  *color.RGBA
}

const (
	BLACK = iota
	BROWN
	RED
	ORANGE
	YELLOW
	GREEN
	BLUE
	VIOLET
	GRAY
	WHITE
	GOLD
	SILVER
)

var (
	Bar_BLACK  = &color.RGBA{0, 0, 0, 1}
	Bar_BROWN  = &color.RGBA{115, 78, 48, 1}
	Bar_RED    = &color.RGBA{255, 0, 0, 1}
	Bar_ORANGE = &color.RGBA{255, 128, 0, 1}
	Bar_YELLOW = &color.RGBA{255, 255, 0, 1}
	Bar_GREEN  = &color.RGBA{0, 255, 0, 1}
	Bar_BLUE   = &color.RGBA{0, 0, 255, 1}
	Bar_VIOLET = &color.RGBA{170, 0, 255, 1}
	Bar_GRAY   = &color.RGBA{105, 105, 105, 1}
	Bar_WHITE  = &color.RGBA{255, 255, 255, 1}
	Bar_GOLD   = &color.RGBA{255, 215, 0, 1}
	Bar_SILVER = &color.RGBA{192, 192, 192, 1}
)
