package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	X, Y, Width, Hight int32
	rectangleBody      rl.Rectangle
	textPostionStart   int
	FuncExce           func()
	BackgroundColor    rl.Color
	TextColor          rl.Color
	FontSize           int32
	Text               string
	Font               rl.Font
	round              float32
	Icon               Icons
	IconSize           int8
	IconColor          rl.Color
	iconVector         rl.Vector2
	app                *App
	hide               bool
	MemberModal        bool
}

func (bu *Button) Draw(x, y, width, hight int32) {
	if !bu.hide {
		bu.body(x, y, width, hight)
	}
}

func (bu *Button) Hide() {
	bu.hide = false
}

func (bu *Button) Show() {
	bu.hide = true
}

func (bu *Button) body(x, y, width, hight int32) {
	bu.rectangleBody.X = float32(x)
	bu.rectangleBody.Y = float32(y)
	bu.rectangleBody.Width = float32(width)
	bu.rectangleBody.Height = float32(hight)
	rl.DrawRectangleRounded(bu.rectangleBody, bu.round, 0, bu.BackgroundColor)
	if !bu.MemberModal {
		rl.BeginScissorMode(x, y, width, hight)
	}
	if bu.Icon.X != 0 {
		bu.app.Icon.Draw(bu.Icon, float32((int32(bu.iconVector.X)+x+(width/2)))-float32(rl.MeasureText(bu.Text, bu.FontSize)/2), float32((int32(bu.iconVector.Y) + y + (hight / 3))), bu.IconSize, bu.IconColor)
	}
	rl.DrawTextEx(bu.Font, bu.Text, rl.Vector2{X: float32((x + (width / 2))) - float32(rl.MeasureText(bu.Text, bu.FontSize)/2), Y: float32((y + (hight / 4)))}, float32(bu.FontSize), 3, bu.TextColor)
	if !bu.MemberModal {
		rl.EndScissorMode()
	}
	bu.Event()
}

func (bu *Button) Event() {
	checkMouseCollisionRec := rl.CheckCollisionRecs(bu.rectangleBody, rl.Rectangle{float32(rl.GetMouseX()), float32(rl.GetMouseY()), 1, 1})
	if checkMouseCollisionRec && bu.isMemberModal() {
		bu.BackgroundColor.A = 160
		rl.DrawRectangleRoundedLines(bu.rectangleBody, bu.round, 0, 2, rl.Blue)
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			checkFuncIsInstince := bu.FuncExce
			if checkFuncIsInstince == nil {
				println("please set function with FuncExce example object.FuncExce = func")
				return
			}
			bu.FuncExce()
		}
	}
	if !checkMouseCollisionRec {
		bu.BackgroundColor.A = 220
	}
}

func (bu *Button) SetPosition(x, y int32) {
	bu.X = x
	bu.Y = y
}

func (bu *Button) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(bu.X), float32(bu.Y), float32(bu.Width), float32(bu.Hight))
}

func (bu *Button) DrawInterface(x, y, width, hight int32) {
	if !bu.hide {
		bu.body(x, y, width, hight)
	}
}

func (bu *Button) isMemberModal() bool {
	if bu.MemberModal {
		return true
	}
	return !bu.app.Lock
}

func (bu *Button) EnableMemberModal() {
	bu.MemberModal = true
}
