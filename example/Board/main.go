package main

import (
	gox "github.com/mohsengreen1388/gox/widgets"
)

func main() {

	app := gox.Init("App", 700, 400)

	// init board
	boardOne := app.NewBoard()
	boardTwo := app.NewBoard()

	//set DefulteBoard is necessary
	app.DefulteBoard = boardOne

	//init label for content board
	bu1 := app.NewButton()
	bu1.Text = "page two"
	bu1.BackgroundColor = gox.PINK
	bu1.FuncExce = func() {
		changeBoardToTwo(app,boardTwo)
	}

	bu2 := app.NewButton()
	bu2.Text = "To page one"
	bu2.FuncExce = func() {
		changeBoardToTwo(app,boardOne)
	}

	boardOne.AddChild(bu1)
	boardTwo.AddChild(bu2)

	gox.Update(app, func() {
		bu1.SetPosition(250,200)
		bu2.SetPosition(250,200)
	})
}

func changeBoardToTwo(a *gox.App,b *gox.Board){
	a.DefulteBoard = b
}

func changeBoardToOne(a *gox.App,b *gox.Board){
	a.DefulteBoard = b
}
