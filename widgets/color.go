package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

var LIGHTGRAY = rl.LightGray
var GRAY = rl.Gray
var DARKGRAY = rl.DarkGray
var YELLOW = rl.Yellow
var GOLD = rl.Gold
var ORANGE = rl.Orange
var PINK = rl.Pink
var Red = rl.Red
var MAROON = rl.Maroon
var GREEN = rl.Green
var LIME = rl.Lime
var DARKGREEN = rl.DarkGreen
var SKYBLUE = rl.SkyBlue
var BLUE = rl.Blue
var DARKBLUE = rl.DarkBlue
var PURPLE = rl.Purple
var VIOLET = rl.Violet
var DARKPURPLE = rl.DarkPurple
var BEIGE = rl.Beige
var BROWN = rl.Brown
var DARKBROWN = rl.DarkBlue
var WHITE = rl.White
var BLACK = rl.Black
var BLANK = rl.Blank
var MAGENTA = rl.Magenta
var RAYWHITE = rl.RayWhite

func CustomColor(red, green, blue, alpha uint8) color.RGBA {
	return color.RGBA{R: red, G: green, B: blue, A: alpha}
}
