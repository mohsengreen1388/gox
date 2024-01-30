package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SPEED = 10
	ZERO  = 0
)

type Board struct {
	Font                                      rl.Font
	slideSize                                 int32
	app                                       *App
	Iteams                                    []InterfaceDraw
	speedScroolUpDown, countUpManyOutOFWindow int32
	checkIsInsideWindowObject                 bool
}

func (win *Board) Draw(x, y, width, hight int32) {
	win.slideSize = win.app.SlideSize
	win.body(win.slideSize+x, y, width, hight)
}

func (win *Board) body(x, y, width, hight int32) {
	for _, iteam := range win.Iteams {
		if (int32(iteam.Coord().Y)+y)+win.speedScroolUpDown > hight-int32(win.Iteams[len(win.Iteams)-1].Coord().Height) {
			win.checkIsInsideWindowObject = true
		}

		iteam.DrawInterface(x+int32(iteam.Coord().X), (int32(iteam.Coord().Y)+y)+win.speedScroolUpDown, int32(iteam.Coord().Width), int32(iteam.Coord().Height))
	}
	win.Event()
}

func (win *Board) Event() {
	if rl.GetMouseWheelMove() < ZERO && win.countUpManyOutOFWindow != ZERO || rl.IsKeyDown(rl.KeyUp) && win.countUpManyOutOFWindow != ZERO {
		win.speedScroolUpDown += SPEED
		win.countUpManyOutOFWindow++
	}
	if rl.GetMouseWheelMove() > ZERO && win.checkIsInsideWindowObject || rl.IsKeyDown(rl.KeyDown) && win.checkIsInsideWindowObject {
		win.speedScroolUpDown -= SPEED
		win.countUpManyOutOFWindow--
		win.checkIsInsideWindowObject = false
	}
}

func (win *Board) AddChild(iteam ...InterfaceDraw) {
	win.Iteams = append(win.Iteams, iteam...)
}
