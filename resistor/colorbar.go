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
	Bar_BLACK  = &color.RGBA{0, 0, 0, 255}
	Bar_BROWN  = &color.RGBA{115, 78, 48, 255}
	Bar_RED    = &color.RGBA{255, 0, 0, 255}
	Bar_ORANGE = &color.RGBA{255, 128, 0, 255}
	Bar_YELLOW = &color.RGBA{255, 255, 0, 255}
	Bar_GREEN  = &color.RGBA{0, 255, 0, 255}
	Bar_BLUE   = &color.RGBA{0, 0, 255, 255}
	Bar_VIOLET = &color.RGBA{170, 0, 255, 255}
	Bar_GRAY   = &color.RGBA{105, 105, 105, 255}
	Bar_WHITE  = &color.RGBA{255, 255, 255, 255}
	Bar_GOLD   = &color.RGBA{255, 215, 0, 255}
	Bar_SILVER = &color.RGBA{192, 192, 192, 255}
)
