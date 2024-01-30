package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuBar struct {
	x, y, width, hight                                      int32
	Font                                                    rl.Font
	FontSize                                                int32
	TextColor                                               rl.Color
	BackgroundColor                                         rl.Color
	MemberModal                                             bool
	app                                                     *App
	menuFloatIteams                                         []*MenuFloat
	nameIteams                                              []string
	selectMenu                                              *MenuFloat
	selectx, selecty                                        int32
	calculatSpace, widthPerIteam, calculatSpaceMinusCurrent float32
	spaceBetwenPerIteam                                     float32
}

func (bar *MenuBar) Draw() {
	bar.app.MenuBarSize = bar.hight
	bar.body(bar.x, bar.y)
	if bar.selectMenu != nil {
		bar.selectMenu.Draw(bar.selectx, bar.selecty)
	}
}

func (bar *MenuBar) body(x, y int32) {
	bar.width = int32(rl.GetScreenWidth())
	rl.DrawRectangle(bar.x, bar.y, bar.width, bar.hight,bar.BackgroundColor)

	bar.calculatSpace = 0
	bar.widthPerIteam = 0
	bar.calculatSpaceMinusCurrent = 0

	for index, value := range bar.nameIteams {
		bar.calculatSpace += rl.MeasureTextEx(bar.Font, bar.nameIteams[index], float32(bar.FontSize), 1).X + bar.spaceBetwenPerIteam
		bar.calculatSpaceMinusCurrent = bar.calculatSpace - rl.MeasureTextEx(bar.Font, bar.nameIteams[index], float32(bar.FontSize), 1).X
		bar.widthPerIteam = rl.MeasureTextEx(bar.Font, bar.nameIteams[index], float32(bar.FontSize), 1).X
		bar.drawMenu(value, bar.menuFloatIteams[index], int32(bar.calculatSpaceMinusCurrent), 0, int32(bar.widthPerIteam))
	}
}

func (bar *MenuBar) event(menuFloat *MenuFloat, x, y, width, hight int32) {

	if rl.CheckCollisionRecs(rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(width), Height: float32(hight)},
		rl.Rectangle{X: float32(rl.GetMouseX()), Y: float32(rl.GetMouseY()), Width: 0, Height: 0}) &&
		rl.CheckCollisionRecs(rl.Rectangle{X: float32(bar.x), Y: float32(bar.y), Width: float32(bar.width), Height: float32(bar.hight)},
			rl.Rectangle{X: float32(rl.GetMouseX()), Y: float32(rl.GetMouseY()), Width: 0, Height: 0}) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			bar.selectMenu = nil
			bar.selectMenu = menuFloat
			bar.selectx = x
			bar.selecty = y + 20
		}
	}

	if bar.selectMenu != nil && !rl.CheckCollisionRecs(rl.Rectangle{X: float32(bar.x), Y: float32(bar.y), Width: float32(bar.width), Height: float32(bar.hight)},
		rl.Rectangle{X: float32(rl.GetMouseX()), Y: float32(rl.GetMouseY()), Width: 0, Height: 0}) &&
		!rl.CheckCollisionRecs(rl.Rectangle{X: float32(bar.selectMenu.X), Y: float32(bar.selectMenu.Y), Width: float32(bar.selectMenu.Width), Height: float32(bar.selectMenu.Hight)},
			rl.Rectangle{X: float32(rl.GetMouseX()), Y: float32(rl.GetMouseY()), Width: 0, Height: 0}) {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			bar.selectMenu = nil
		}
	}
}

func (bar *MenuBar) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(bar.x), float32(bar.y), float32(bar.width), float32(bar.hight))
}

func (bar *MenuBar) DrawInterface(x, y, width, hight int32) {
	bar.body(bar.x, bar.y)
}

func (bar *MenuBar) EnableMemberModal() {
	bar.MemberModal = true
}

func (bar *MenuBar) AddItem(name string, menuFloat *MenuFloat) {
	bar.nameIteams = append(bar.nameIteams, name)
	bar.menuFloatIteams = append(bar.menuFloatIteams, menuFloat)
}

func (bar *MenuBar) drawMenu(name string, menu *MenuFloat, x, y, w int32) {
	rl.BeginScissorMode(bar.x, bar.y, bar.width, bar.hight)
	rl.DrawTextEx(bar.Font, name, rl.Vector2{float32(x), float32(y)}, float32(bar.FontSize), 1, bar.TextColor)
	rl.EndScissorMode()
	bar.event(menu, x, y, w, 20)
}
