package gox

import (
	//"strconv"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Progress struct {
	X, Y, Width, Hight int32
	Font               rl.Font
	MemberModal        bool
	Color              rl.Color
	app                *App
	Value              int32
	Auto               bool
	ShowPrecentage     bool
	percentageTostring string
	countProgres       float32
	precentage         float32
}

func (pro *Progress) Draw(x, y float32) {
	pro.X = int32(x)
	pro.Y = int32(y)
	pro.body(x, y)
}

func (pro *Progress) body(x, y float32) {
	black := rl.Black
	black.A = 20
	pro.countProgres = pro.calculatePrecentage()
	pro.moveProgress()
	rl.DrawRectangleRounded(rl.Rectangle{x, y, float32(pro.Width), float32(pro.Hight)}, 0.4, 0, black)
	rl.DrawRectangleRounded(rl.Rectangle{x, y, float32(pro.countProgres), float32(pro.Hight)}, 0.4, 0, pro.Color)
	pro.percentageTostring = strconv.FormatInt(int64(pro.Value), 10)
	if pro.ShowPrecentage {
		rl.DrawTextEx(pro.Font, (pro.percentageTostring)+"%", rl.Vector2{x + 50, y + 2}, 20, 1, rl.Black)
	}
}

func (pro *Progress) SetPosition(x, y int32) {
	pro.X = x
	pro.Y = y
}

func (pro *Progress) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(pro.X), float32(pro.Y), float32(pro.Width), float32(pro.Hight))
}

func (pro *Progress) DrawInterface(x, y, width, hight int32) {
	pro.body(float32(x), float32(y))
}

func (pro *Progress) EnableMemberModal() {
	pro.MemberModal = true
}

func (pro *Progress) moveProgress() {
	if pro.Auto {
		if pro.Value == 100 {
			pro.Value = 0
		}
		pro.Value++
	}
}

func (pro *Progress) calculatePrecentage() float32 {
	pro.precentage = float32(((pro.Value * pro.Width) / 100))
	return pro.precentage
}
