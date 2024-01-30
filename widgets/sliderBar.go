package gox

import (
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SlideBar struct {
	X, Y, Width, Hight  int32
	Font                rl.Font
	MemberModal         bool
	Color               rl.Color
	app                 *App
	Value               int32
	Auto                bool
	ShowPrecentage      bool
	percentageTostring  string
	countProgres        float32
	precentage          float32
	circleRange         rl.Vector2
	CircleRangeColor    rl.Color
	PercentageColor     rl.Color
	TextPercentageColor rl.Color
	textPercentageSize  int32
	round               int32
	Min, Max            int32
}

func (sl *SlideBar) Draw(x, y float32) {
	sl.X = int32(x)
	sl.Y = int32(y)
	sl.body(x, y)
}

func (sl *SlideBar) body(x, y float32) {
	black := rl.Black
	black.A = 20
	sl.round = 12
	sl.countProgres = float32(sl.Value) * sl.calculatePrecentage()

	sl.circleRange = rl.Vector2{X: float32(x) + float32((sl.countProgres)), Y: float32((y + 12))}
	rl.DrawRectangleRounded(rl.Rectangle{x, y, float32(sl.Width), float32(sl.Hight)}, 0.4, 0, black)
	rl.DrawRectangleRounded(rl.Rectangle{x, y, float32(sl.countProgres), float32(sl.Hight)}, 0.4, 0, sl.Color)
	rl.DrawCircle(int32(sl.circleRange.X), int32(sl.circleRange.Y), 12, sl.CircleRangeColor)

	if sl.ShowPrecentage {
		sl.percentageTostring = strconv.FormatInt(int64(sl.Value), 10)
		rl.DrawTextEx(sl.Font, (sl.percentageTostring), rl.Vector2{x+(float32(sl.Width)/2.5), y+5}, float32(sl.textPercentageSize), 1, sl.TextPercentageColor)
	}

	sl.moveProgress(int32(x), int32(y))
}

func (sl *SlideBar) SetPosition(x, y int32) {
	sl.X = x
	sl.Y = y
}

func (sl *SlideBar) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(sl.X), float32(sl.Y), float32(sl.Width), float32(sl.Hight))
}

func (sl *SlideBar) DrawInterface(x, y, width, hight int32) {
	sl.body(float32(x), float32(y))
}

func (sl *SlideBar) isMemberModal() bool {
	if sl.MemberModal {
		return true
	}
	return !sl.app.Lock
}

func (sl *SlideBar) EnableMemberModal() {
	sl.MemberModal = true
}

func (sl *SlideBar) moveProgress(x, y int32) {

	CheckCollisionCircles := rl.CheckCollisionCircles(rl.Vector2{sl.circleRange.X, sl.circleRange.Y}, float32(sl.round), rl.Vector2{float32(sl.app.MouseX()), float32(sl.app.MouseY())}, float32(0))
	CheckCollisionRecs := rl.CheckCollisionCircleRec(rl.Vector2{float32(rl.GetMouseX()), float32(rl.GetMouseY())}, float32(0), rl.Rectangle{float32(x), float32(y), float32(sl.Width), float32(sl.Hight)})

	if (CheckCollisionCircles || CheckCollisionRecs) && rl.IsMouseButtonDown(rl.MouseLeftButton) && sl.isMemberModal() {
		if rl.GetMouseDelta().X > 0 {
			if sl.Value < sl.Max {
				sl.countProgres += sl.calculatePrecentage()
				sl.Value++
			}
		}

		if rl.GetMouseDelta().X < 0 {
			if sl.Value > sl.Min {
				sl.countProgres -= sl.calculatePrecentage()
				sl.Value--
			}
		}
	}

}

func (sl *SlideBar) calculatePrecentage() float32 {
	sl.precentage = float32(sl.Width) / float32(sl.Max)
	return sl.precentage
}
