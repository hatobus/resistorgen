package testing

import (
	"reflect"
	"testing"

	"github.com/hatobus/resistorgen/resistor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenColor(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	testcase := []struct {
		Name      string
		RecValue  string
		ColorBand *resistor.Colorbar
	}{
		{
			Name:     "1kohm",
			RecValue: "1K",
			ColorBand: &resistor.Colorbar{
				FirstBand:  resistor.Bar_BROWN,
				SecondBand: resistor.Bar_BLACK,
				ThirdBand:  resistor.Bar_RED,
				Tolerance:  resistor.Bar_GOLD,
			},
		},
		{
			Name:     "470ohm",
			RecValue: "470",
			ColorBand: &resistor.Colorbar{
				FirstBand:  resistor.Bar_YELLOW,
				SecondBand: resistor.Bar_VIOLET,
				ThirdBand:  resistor.Bar_BROWN,
				Tolerance:  resistor.Bar_GOLD,
			},
		},
		{
			Name:     "4k7ohm",
			RecValue: "4k7",
			ColorBand: &resistor.Colorbar{
				FirstBand:  resistor.Bar_YELLOW,
				SecondBand: resistor.Bar_VIOLET,
				ThirdBand:  resistor.Bar_RED,
				Tolerance:  resistor.Bar_GOLD,
			},
		},
		{
			Name:     "47kohm",
			RecValue: "47k",
			ColorBand: &resistor.Colorbar{
				FirstBand:  resistor.Bar_YELLOW,
				SecondBand: resistor.Bar_VIOLET,
				ThirdBand:  resistor.Bar_ORANGE,
				Tolerance:  resistor.Bar_GOLD,
			},
		},
	}

	for _, tc := range testcase {
		t.Run(tc.Name, func(t *testing.T) {
			color, err := resistor.GenerateColor(tc.RecValue)
			assert.NoError(err)

			require.True(reflect.DeepEqual(tc.ColorBand, color))
		})
	}

}
