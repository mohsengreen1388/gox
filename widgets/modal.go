package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Modal struct {
	X, Y, Width, Hight int32
	Font               rl.Font
	app                *App
	Iteams             []InterfaceDraw
	Rounde             float32
	BackgroundColor    rl.Color
	close              bool
	shadowColor        rl.Color
}

func (mo *Modal) Draw(x, y, width, hight int32) {
	mo.X = x
	mo.Y = y
	mo.Width = width
	mo.Hight = hight
	if !mo.close {
		mo.body(x, y, width, hight)
	}
}

func (mo *Modal) body(x, y, width, hight int32) {
	mo.shadowColor.A = 30
	rl.DrawRectangleRounded(rl.Rectangle{X: float32(x - 5), Y: float32(y - 2), Width: float32(width + 10), Height: float32(hight + 4)}, mo.Rounde, 0, mo.shadowColor)
	rl.DrawRectangleRounded(rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(width), Height: float32(hight)}, mo.Rounde, 0, mo.BackgroundColor)
	for _, iteam := range mo.Iteams {
		rl.BeginScissorMode(x, y, width, hight)
		iteam.EnableMemberModal()
		iteam.DrawInterface(x+int32(iteam.Coord().X), (int32(iteam.Coord().Y) + y), int32(iteam.Coord().Width), int32(iteam.Coord().Height))
		rl.EndScissorMode()
	}
	mo.event()
}

func (mo *Modal) event() {
	check := rl.CheckCollisionRecs(rl.Rectangle{float32(mo.X), float32(mo.Y), float32(mo.Width), float32(mo.Hight)}, rl.Rectangle{float32(rl.GetMouseX()), float32(rl.GetMouseY()), 0, 0})

	if check {
		mo.app.Lock = true
	} else {
		mo.app.Lock = false
	}
}

func (mo *Modal) SetPosition(x, y, width, hight int32) {
	mo.X = x
	mo.Y = y
	mo.Width = width
	mo.Hight = hight
}

func (mo *Modal) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(mo.X), float32(mo.Y), float32(mo.Width), float32(mo.Hight))
}

func (mo *Modal) DrawInterface(x, y, width, hight int32) {
	if !mo.close {
		mo.body(x, y, width, hight)
	}
}

func (mo *Modal) Close() {
	mo.close = true
}

func (mo *Modal) Open() {
	mo.close = false
}

func (mo *Modal) EnableMemberModal() {}

func (mo *Modal) AddChild(iteam ...InterfaceDraw) {
	mo.Iteams = append(mo.Iteams, iteam...)
}
