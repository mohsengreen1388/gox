package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Label struct {
	X, Y, Width, Hight int32
	Font               rl.Font
	TextColor          rl.Color
	TextSize           float32
	Textspace          float32
	MemberModal        bool
	app                *App
	Text               string
	Icon               Icons
	IconColor          rl.Color
	iconVector         rl.Vector2
}

func (la *Label) Draw(text string, x, y float32) {
	la.X = int32(x)
	la.Y = int32(y)
	la.Text = text
	la.body(la.Text, x, y)
}

func (la *Label) body(text string, x, y float32) {
	if la.Icon.X != 0 {
		la.app.Icon.Draw(la.Icon, x+la.iconVector.X-float32(la.TextSize), y+la.iconVector.Y, int8(la.TextSize), la.IconColor)
	}
	rl.DrawTextEx(la.Font, text, rl.Vector2{x, y}, la.TextSize, la.Textspace, la.TextColor)
}

func (la *Label) SetPosition(x, y int32) {
	la.X = x
	la.Y = y
}

func (la *Label) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(la.X), float32(la.Y), float32(la.Width), float32(la.Hight))
}

func (la *Label) DrawInterface(x, y, width, hight int32) {
	la.body(la.Text, float32(x), float32(y))
}

func (la *Label) EnableMemberModal() {
	la.MemberModal = true
}
