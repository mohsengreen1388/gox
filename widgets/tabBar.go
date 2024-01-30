package gox

import (
	"math"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type TabBar struct {
	X, Y, Width, Hight int32
	Font               rl.Font
	app                *App
	ChildBoard         bool
	Count              int8
	Buttons            []*Button
	AutoSize           bool
	Minsize            int32
	MemberModal        bool
	totalMinsize       int32
	SpeedScroll        int32
	step               int32
	FontSize           int32
	TextColor          rl.Color
	BackgroundColor    rl.Color
}

func (tab *TabBar) Draw(x, y int32) {
	tab.X = x
	tab.Y = y
	tab.body(x, y)
}

func (tab *TabBar) body(x, y int32) {
	for index, bu := range tab.Buttons {
		bu.X = x + tab.X + int32(tab.getSize())*int32(index) + tab.step
		bu.Y = y
		bu.Width = int32(tab.getSize())
		bu.Hight = tab.Hight
		bu.Draw(bu.X, bu.Y, bu.Width, bu.Hight)
	}
	tab.event()
}

func (tab *TabBar) AddChildButtons() {
	for i := 0; i < int(tab.Count); i++ {
		bu := tab.app.NewButton()
		bu.Text = "iteam" + strconv.FormatInt(int64(i), 10)
		bu.round = 0
		bu.TextColor = tab.TextColor
		bu.FontSize = tab.FontSize
		bu.BackgroundColor = tab.BackgroundColor
		tab.Buttons = append(tab.Buttons, bu)
	}
}

func (tab *TabBar) SetPosition(x, y int32) {
	tab.X = x
	tab.Y = y
}

func (tab *TabBar) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(tab.X), float32(tab.Y), float32(tab.Width), float32(tab.Hight))
}

func (tab *TabBar) DrawInterface(x, y, width, hight int32) {
	tab.body(x, y)
}

func (tab *TabBar) getSize() int {
	if tab.AutoSize {
		sizEachButton := tab.childBoard() / int32(tab.Count)
		if sizEachButton < tab.Minsize {
			return int((sizEachButton) + int32(tab.Minsize))
		}
		tab.step = 0
		return int(sizEachButton + 1)
	}
	return int(tab.Width)
}

func (tab *TabBar) childBoard() int32 {
	if tab.ChildBoard {
		return tab.app.GetScreenWidth()
	}
	return int32(rl.GetScreenWidth())
}

func (tab *TabBar) AddItemFun(item int, f func()) {
	tab.Buttons[item].FuncExce = f
}
func (tab *TabBar) ItemName(item int, name string) {
	tab.Buttons[item].Text = name
}
func (tab *TabBar) event() {
	tab.totalMinsize = tab.Minsize * int32(tab.Count)
	CheckCollisionWithMouse := rl.CheckCollisionRecs(rl.Rectangle{float32(tab.X), float32(tab.Y), float32(tab.totalMinsize), float32(tab.Hight)},
		rl.Rectangle{float32(tab.app.MouseX()), float32(tab.app.MouseY()), 0, 0})

	if CheckCollisionWithMouse && rl.GetMouseDelta().X > 0 && tab.step < 0 {
		tab.step += tab.SpeedScroll
	}
	if CheckCollisionWithMouse && rl.GetMouseDelta().X < 0 && int32(math.Abs(float64(tab.step))) < tab.totalMinsize-int32(tab.Count/2) {
		tab.step -= tab.SpeedScroll
	}
}

func (tab *TabBar) EnableMemberModal() {
	tab.MemberModal = true
}

func (tab *TabBar) Unload() {

}
