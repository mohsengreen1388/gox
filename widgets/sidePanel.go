package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction int

const (
	Right Direction = iota
	Left
)

type SidePanel struct {
	X, Y, Width, Hight int32
	Font               rl.Font
	app                *App
	model              *Modal
	Iteams             []InterfaceDraw
	Rounde             float32
	BackgroundColor    rl.Color
	shadowColor        rl.Color
	Direction          Direction
}

func (si *SidePanel) Draw() {
	si.X = si.caclculatDirection()
	si.Width = si.Width
	si.body(si.X, si.Y, si.Width, si.Hight)
}

func (si *SidePanel) body(x, y, width, hight int32) {
	si.model.app = si.app
	si.model.Font = si.Font
	si.model.Rounde = si.Rounde
	si.model.BackgroundColor = si.BackgroundColor
	si.model.shadowColor = si.shadowColor

	if si.Direction == Right && si.app.SlideSize > 0 {

	} else {
		si.model.Draw(x, y, width, hight)
	}
}

func (si *SidePanel) Coord() rl.Rectangle {
	return si.model.Coord()
}

func (si *SidePanel) DrawInterface(x, y, width, hight int32) {
	si.body(x, y, width, hight)
}

func (si *SidePanel) AddChild(iteam ...InterfaceDraw) {
	si.model.Iteams = append(si.model.Iteams, iteam...)
}

func (si *SidePanel) caclculatDirection() int32 {
	if si.Direction == Right {
		return 5
	}
	return int32(rl.GetScreenWidth()) - si.Width - 5
}
