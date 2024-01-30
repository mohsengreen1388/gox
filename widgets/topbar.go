package gox

import rl "github.com/gen2brain/raylib-go/raylib"

type TopBar struct {
	FuncExce       func()
	ColorBacground rl.Color
	TextColor      rl.Color
	FontSize       int32
	Font           rl.Font
	Text           string
	App            *App
	Button         *Button
	slideEnable    bool
	slideRectangle rl.Rectangle
	rectangleMouse rl.Rectangle
	IconTop        Icons
	IconColor      rl.Color
	iconSize       int8
	TextLogo       string
	TextLogoSize   float32
	TextLogoColor  rl.Color
	fixed          bool
	ImageWidth     float32
	Imagehight     float32
	Buttons        []*Button
	app            *App
}

func (top *TopBar) Draw() {
	top.body()
}

func (top *TopBar) body() {

	if !top.fixed {
		rl.DrawRectangle(0, 0+top.app.MenuBarSize, int32(rl.GetScreenWidth()), 50, rl.SkyBlue)
		rl.DrawTextEx(top.Font, top.Text, rl.Vector2{X: float32(rl.GetScreenWidth()/2) - float32(rl.MeasureText(top.Text, top.FontSize)/2), Y: 10 + float32(top.app.MenuBarSize)}, float32(top.FontSize), 3, top.TextColor)
		top.Button.Width = 50
		top.Button.Hight = 50
		top.Button.Draw(0, 0+top.app.MenuBarSize, 50, 50)
		top.Button.FuncExce = top.Open
		top.Button.round = 0
		top.Button.Text = ""
	} else {
		if top.slideEnable {
			top.slideEnable = true
		}
	}

	top.closeOrOpenSlide()
}

// only work when topbar be fixed
func (top *TopBar) Open() {
	top.slideEnable = true
}

func (top *TopBar) Close() {
	top.app.SlideSize = 0
	top.slideEnable = false
}

func (top *TopBar) Fixed() {
	top.fixed = true
	top.slideEnable = true
}

var color1 rl.Color

func (top *TopBar) closeOrOpenSlide() {
	if top.slideEnable {
		screenWidth := float32(rl.GetScreenWidth()) / float32(top.calculateSize())
		ScreenHight := rl.GetScreenHeight()
		top.app.SlideSize = int32(screenWidth)

		top.slideRectangle = rl.Rectangle{X: 0, Y: 0, Width: float32(screenWidth), Height: float32(ScreenHight)}
		top.rectangleMouse = rl.Rectangle{X: float32(rl.GetMouseX()), Y: float32(rl.GetMouseY()), Width: 0, Height: 0}

		rl.BeginScissorMode(0, 0, int32(screenWidth), int32(ScreenHight))
		rl.DrawRectangle(int32(top.slideRectangle.X), int32(top.slideRectangle.Y),
			int32(top.slideRectangle.Width), int32(top.slideRectangle.Height),
			rl.RayWhite)
		top.app.Icon.Draw(top.IconTop, screenWidth/3, 20, top.iconSize, top.IconColor)
		rl.DrawTextEx(top.Font, top.TextLogo, rl.Vector2{screenWidth / 4, top.Imagehight}, top.TextLogoSize, 1, top.TextLogoColor)
		rl.DrawLine(0, int32(top.Imagehight)+50, int32(screenWidth), int32(top.Imagehight)+50, rl.Black)

		for index, bu := range top.Buttons {
			bu.X = 5
			bu.Y = int32(top.Imagehight) + int32((55*index)+60)
			bu.Width = int32(screenWidth) - 10
			bu.Hight = 46
			bu.Draw(bu.X, bu.Y, bu.Width, bu.Hight)
		}
		rl.EndScissorMode()

		if !rl.CheckCollisionRecs(top.slideRectangle, top.rectangleMouse) && !top.fixed {
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				top.Close()
			}
		}

	}
}

func (top *TopBar) Event() {

}

func (top *TopBar) calculateSize() float32 {

	if rl.GetScreenWidth() < 800 && rl.GetScreenWidth() > 600 {
		return 3
	}
	if rl.GetScreenWidth() < 600 {
		return 2.7
	}

	return 4.7
}

func (top *TopBar) AddChildButtons(iteam ...*Button) {
	top.Buttons = append(top.Buttons, iteam...)
}
func (top *TopBar) Unload() {

}
