package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CheckBox struct {
	X, Y, Width, Hight int32
	Font               rl.Font
	MemberModal        bool
	BackgroundColor    rl.Color
	CircleColor        rl.Color
	circlePos          rl.Vector2
	Check              bool
	app                *App
}

func (ch *CheckBox) Draw(x, y float32) {
	ch.X = int32(x)
	ch.Y = int32(y)
	ch.body(x, y)
}

func (ch *CheckBox) body(x, y float32) {
	rl.DrawRectangleRounded(rl.Rectangle{x, y, float32(ch.Width), float32(ch.Hight)}, 2, 0, ch.BackgroundColor)
	ch.CircleColor.A = 200
	rl.DrawPoly(rl.Vector2{X: x + ch.circlePos.X + 10, Y: y + ch.circlePos.Y + 12}, 20, 14, 0, ch.CircleColor)
	ch.Event(x, y)
}

func (ch *CheckBox) Event(x, y float32) {
	checkcoll := rl.CheckCollisionRecs(rl.Rectangle{x, y, float32(ch.Width), float32(ch.Hight)}, rl.Rectangle{float32(ch.app.MouseX()), float32(ch.app.MouseY()), 0, 0})

	if checkcoll && !ch.Check && rl.IsMouseButtonPressed(rl.MouseLeftButton) && ch.isMemberModal() {
		ch.circlePos.X = 30
		ch.BackgroundColor = rl.Green
		ch.Check = true
	} else {
		if checkcoll && ch.Check && rl.IsMouseButtonPressed(rl.MouseLeftButton) && ch.isMemberModal() {
			black := rl.Black
			black.A = 80
			ch.circlePos.X = 0
			ch.BackgroundColor = black
			ch.Check = false
		}
	}

}

func (ch *CheckBox) SetPosition(x, y int32) {
	ch.X = x
	ch.Y = y
}

func (ch *CheckBox) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(ch.X), float32(ch.Y), float32(ch.Width), float32(ch.Hight))
}


func (ch *CheckBox) DrawInterface(x, y, width, hight int32) {
	ch.body(float32(x), float32(y))
}

func (ch *CheckBox) isMemberModal() bool {
	if ch.MemberModal {
		return true
	}
	return !ch.app.Lock
}
func (ch *CheckBox) EnableMemberModal() {
	ch.MemberModal = true
}
