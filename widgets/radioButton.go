package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RadioButton struct {
	X, Y, Width, Hight int32
	Font               rl.Font
	MemberModal        bool
	BackgroundColor    rl.Color
	CircleColor        rl.Color
	CircleColorChecked rl.Color
	TextColor          rl.Color
	radius             int8
	Check              bool
	label              Label
	Value              interface{}
	Name               string
	Id                 string
	app                *App
}

func (ra *RadioButton) Draw(x, y float32) {
	ra.X = int32(x)
	ra.Y = int32(y)
	ra.body(x, y)
}

func (ra *RadioButton) body(x, y float32) {
	rl.DrawCircleLines(int32(x), int32(y), float32(ra.radius), ra.BackgroundColor)
	rl.DrawPoly(rl.Vector2{X: x, Y: y}, 20, float32(ra.radius)-3, 1, ra.CircleColor)
	ra.label.TextColor = ra.TextColor
	ra.label.Draw(ra.Name, x+20, y-(ra.label.TextSize/2))
	ra.event(x, y)
}

func (ra *RadioButton) event(x, y float32) {
	checkcoll := rl.CheckCollisionPointCircle(rl.Vector2{X: float32(ra.app.MouseX()), Y: float32(ra.app.MouseY())}, rl.Vector2{X: x, Y: y}, float32(ra.radius))

	if checkcoll && !ra.Check && rl.IsMouseButtonPressed(rl.MouseLeftButton) && ra.isMemberModal() {
		ra.Checked()
		ra.UnCheckedOther()
	} else {
		if checkcoll && ra.Check && rl.IsMouseButtonPressed(rl.MouseLeftButton) && ra.isMemberModal() {
			ra.UnChecked()
		}
	}

}

func (ra *RadioButton) SetPosition(x, y int32) {
	ra.X = x
	ra.Y = y
}

func (ra *RadioButton) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(ra.X), float32(ra.Y), float32(ra.Width), float32(ra.Hight))
}

func (ra *RadioButton) DrawInterface(x, y, width, hight int32) {
	ra.body(float32(x), float32(y))
}

func (ra *RadioButton) isMemberModal() bool {
	if ra.MemberModal {
		return true
	}
	return !ra.app.Lock
}
func (ra *RadioButton) EnableMemberModal() {
	ra.MemberModal = true
}

func (ra *RadioButton) Checked() {
	ra.BackgroundColor.A = 170
	ra.CircleColor = ra.CircleColorChecked
	ra.Check = true
}

func (ra *RadioButton) UnChecked() {
	ra.BackgroundColor.A = 80
	ra.CircleColor = ra.BackgroundColor
	ra.Check = false
}

func (ra *RadioButton) UnCheckedOther() {
	for name, radioptr := range ra.app.RadioButtonChecker {
		if name == radioptr.Name+ra.Id && radioptr != ra {
			radioptr.UnChecked()
		}
	}
}
