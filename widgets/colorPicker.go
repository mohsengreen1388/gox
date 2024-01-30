package gox

import (
	"image/color"

	"github.com/ncruces/zenity"
)

type ColorPicker struct{}

func (co *ColorPicker) SelectColor() (color.Color, error) {
	return zenity.SelectColor()
}
