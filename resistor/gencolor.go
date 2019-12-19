package resistor

import (
	"errors"
	"image/color"
	"strconv"
	"strings"
)

var mul = []string{"k", "K", "m", "M"}

func GenerateColor(colorstr string) (*Colorbar, error) {
	contain, multipler, err := getcontains(colorstr)
	if err != nil {
		return nil, err
	}

	var mag int64
	var firstband, secondband int64

	if multipler == "k" || multipler == "K" {
		mag = 2
	} else if multipler == "m" || multipler == "M" {
		mag = 5
	}

	// finding section for resistor value
	if len(contain[0]) == 1 {
		if contain[1] == "" {
			mag -= 1
			fv, _ := strconv.Atoi(contain[0])
			firstband = int64(fv)
			secondband = 0
		}
	} else if len(contain[0]) == 2 {
		fv, err := strconv.Atoi(contain[0])
		if err != nil {
			return nil, err
		}
		firstband = int64(fv / 10)
		secondband = int64(fv % 10)
	} else {
		// 2のときは xxKΩ や xxΩ として決定しているので大丈夫
		resistor_topvalue, err := strconv.Atoi(contain[0])
		if err != nil {
			return nil, err
		}
		for i := 0; ; i++ {
			last1digit := resistor_topvalue % 10
			resistor_topvalue /= 10
			mag += 1
			if int64(resistor_topvalue/10) == 0 {
				firstband = int64(resistor_topvalue)
				secondband = int64(last1digit)
				break
			}
		}
	}

	fb, err := SolvColorBand(firstband)
	if err != nil {
		return nil, err
	}
	sb, err := SolvColorBand(secondband)
	if err != nil {
		return nil, err
	}
	tb, err := SolvColorBand(mag)
	if err != nil {
		return nil, err
	}

	Bar := &Colorbar{
		FirstBand:  fb,
		SecondBand: sb,
		ThirdBand:  tb,
		Tolerance:  Bar_GOLD,
	}

	return Bar, nil
}

func getcontains(colorstr string) ([]string, string, error) {
	var mulc string
	contains := make([]string, 0, 3)

	val, err := strconv.Atoi(colorstr)
	if err != nil {
		for _, m := range mul {
			contains = strings.Split(colorstr, m)
			if contains[0] == colorstr {
				continue
			} else {
				mulc = m
				break
			}
		}
	} else {
		return []string{string(val)}, "", nil
	}

	if mulc == "" {
		return nil, "", errors.New("invalid string format")
	}

	return contains, mulc, nil
}

func SolvColorBand(num int64) (*color.RGBA, error) {
	switch num {
	case BLACK:
		return Bar_BLACK, nil
	case BROWN:
		return Bar_BROWN, nil
	case RED:
		return Bar_RED, nil
	case ORANGE:
		return Bar_ORANGE, nil
	case YELLOW:
		return Bar_YELLOW, nil
	case GREEN:
		return Bar_GREEN, nil
	case BLUE:
		return Bar_BLUE, nil
	case VIOLET:
		return Bar_VIOLET, nil
	case GRAY:
		return Bar_GRAY, nil
	case WHITE:
		return Bar_WHITE, nil
	case GOLD:
		return Bar_GOLD, nil
	case SILVER:
		return Bar_SILVER, nil
	default:
		return nil, errors.New("Invalid color code")
	}
}
