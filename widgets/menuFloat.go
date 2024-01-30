package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuFloat struct {
	X, Y, Width, Hight  int32
	BackgroundColor     rl.Color
	LineColor           rl.Color
	model               *Modal
	label               *Label
	labelWithoutIcon    *Label
	labelMainIteamIcon  *Label
	labelShortKey       *Label
	offset              int8
	funcAddMenu         []func(x, y int32)
	HightPerIteam       int32
	TotalHight, offsety int32
	selectMain          *MenuFloat
	Enable              bool
}

func (me *MenuFloat) Draw(x, y int32) {
	me.X = x
	me.Y = y
	if me.Enable {
		me.body(x, y)
	}
}

func (me *MenuFloat) body(x, y int32) {

	me.model.Draw(x, y, me.Width, me.Hight)
	me.Hight = int32(me.HightPerIteam) * int32(len(me.funcAddMenu))

	for index, iteam := range me.funcAddMenu {
		iteam(x, me.HightPerIteam*int32(index))
	}

	if me.selectMain != nil {
		me.selectMain.Draw(x+me.selectMain.Width, me.offsety)
	}
}

func (me *MenuFloat) event(x, y int32, Draw func(), mainMenu bool) {
	if me.checkCollisionByMouse(x, y) {
		if !mainMenu {
			me.selectMain = nil
			if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
				Draw()
			}
		}
		if mainMenu {
			me.selectMain = nil
			Draw()
		}

	}
}

func (me *MenuFloat) SetPosition(x, y int32) {
	me.X = x
	me.Y = y
}

func (me *MenuFloat) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(me.X), float32(me.Y), float32(me.Width), float32(me.Hight))
}

func (me *MenuFloat) DrawInterface(x, y, width, hight int32) {
	if me.Enable {
		me.body(x, y)
	}
}

func (me *MenuFloat) Hide() {
	me.Enable = true
}

func (me *MenuFloat) Show() {
	me.Enable = false
}

func (me *MenuFloat) EnableMemberModal() {}

func (me *MenuFloat) addMainIteam(name string, icon Icons, y int32, BackgroundColorCollision rl.Color, function func()) {

	me.label.Icon = icon

	if me.checkCollisionByMouse(me.X, y+me.Y) {
		BackgroundColorCollision = rl.Gray
		BackgroundColorCollision.A = 120
	}

	rl.DrawRectangle(me.X, y+me.Y, me.Width, me.HightPerIteam, BackgroundColorCollision)
	if me.checkLabelIcon(icon) {
		me.label.Draw(name, float32(me.X+int32(me.offset*2)), float32(y+me.Y+5))
	} else {
		me.labelWithoutIcon.Draw(name, float32(me.X+int32(me.offset+5)), float32(y+me.Y+5))
	}
	me.labelMainIteamIcon.Draw("", float32(me.X+me.Width), float32(y+me.Y+5))
	me.label.IconColor = rl.Black
	rl.DrawLine(me.X, y+me.Y+me.HightPerIteam, me.X+me.Width, y+me.Y+me.HightPerIteam, me.LineColor)
	me.event(me.X, y+me.Y, function, true)

}

func (me *MenuFloat) addIteam(name string, icon Icons, y int32, shortkey string, BackgroundColorCollision rl.Color, function func()) {

	me.label.Icon = icon

	if me.checkCollisionByMouse(me.X, y+me.Y) {
		BackgroundColorCollision = rl.Gray
		BackgroundColorCollision.A = 120

	}

	rl.DrawRectangle(me.X, y+me.Y, me.Width, me.HightPerIteam, BackgroundColorCollision)
	if me.checkLabelIcon(icon) {
		me.label.Draw(name, float32(me.X+int32(me.offset*2)), float32(y+me.Y+5))
	} else {
		me.labelWithoutIcon.Draw(name, float32(me.X+int32(me.offset+5)), float32(y+me.Y+5))
	}
	me.labelShortKey.TextSize = 16
	me.labelShortKey.TextColor = rl.Gray
	me.labelShortKey.Draw(shortkey, float32(me.X+me.Width)-float32(rl.MeasureText(shortkey, int32(me.labelShortKey.TextSize)+3)), float32(y+me.Y+7))
	me.label.IconColor = rl.Black
	rl.DrawLine(me.X, y+me.Y+me.HightPerIteam, me.X+me.Width, y+me.Y+me.HightPerIteam, me.LineColor)
	me.event(me.X, y+me.Y, function, false)

}

func (me *MenuFloat) checkCollisionByMouse(x, y int32) bool {
	if rl.CheckCollisionRecs(rl.Rectangle{X: float32(x), Y: float32(y), Width: float32(me.Width), Height: float32(me.HightPerIteam)},
		rl.Rectangle{X: float32(rl.GetMouseX()), Y: float32(rl.GetMouseY()), Width: 0, Height: 0}) {
		return true
	}
	return false
}

func (me *MenuFloat) AddMainMenu(nameItam string, iconItam Icons, menu *MenuFloat) {
	me.funcAddMenu = append(me.funcAddMenu, func(x int32, y int32) {
		me.addMainIteam(nameItam, iconItam, y, me.BackgroundColor, func() {
			if menu != me {
				if me.selectMain == nil {
					me.selectMain = menu
				}
				if me.selectMain == menu {
					me.offsety = me.Y + y
					me.TotalHight = int32(len(menu.funcAddMenu))
				}
			}
		})
	})
}

func (me *MenuFloat) AddSubMenu(nameItam string, iconItam Icons, shortKey string, exeRun func()) {
	me.funcAddMenu = append(me.funcAddMenu, func(x int32, y int32) {
		me.addIteam(nameItam, iconItam, y, shortKey, me.BackgroundColor, func() {
			exeRun()
		})
	})
}

func (me *MenuFloat) checkLabelIcon(icon Icons) bool {
	if icon.X != 0 {
		return true
	}
	return false
}
