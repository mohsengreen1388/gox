package main

import (
	g "github.com/mohsengreen1388/gox/widgets"
	"strings"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var total []string

func main() {
	app := g.Init("Calculator", 300, 330)
	app.Font = app.LoadFont("./Ubuntu-B.ttf", 64)

	rl.ClearWindowState(rl.FlagWindowResizable)
	text := app.NewInput()
	text.X = 10
	text.Y = 20
	text.PlaceholderDefulte = ""
	text.Placeholder = ""
	text.FontSize = 25
	text.ReadOnly = true
	text.Background = rl.DarkGray
	text.TextColor = rl.White
	app.BackgroundColor = rl.Black
	bu := app.NewButton()
	bu.Text = "1"
	bu.FuncExce = func() {
		setNumber(bu.Text)
		justString := strings.Join(total,"")
		text.SetText(justString)
	}
	
	bu2 := app.NewButton()
	bu2.Text = "2"
	bu2.FuncExce = func() {
		setNumber(bu2.Text)
		justString := strings.Join(total,"")
		text.SetText(justString)
	}
	
	bu3 := app.NewButton()
	bu3.Text = "3"
	bu3.FuncExce = func() {
		setNumber(bu3.Text)
		justString := strings.Join(total,"")
		text.SetText(justString)
	}
	
	buplus := app.NewButton()
	buplus.Text = "+"
	buplus.FuncExce = func() {
		setNumber(buplus.Text)
		justString := strings.Join(total,"")
		text.SetText(justString)
	}
	
	bu4 := app.NewButton()
	bu4.Text = "4"

	bu5 := app.NewButton()
	bu5.Text = "5"
	
	bu6 := app.NewButton()
	bu6.Text = "6"
	
	buMinus := app.NewButton()
	buMinus.Text = "-"
	
	bu7 := app.NewButton()
	bu7.Text = "7"

	bu8 := app.NewButton()
	bu8.Text = "8"
	
	bu9 := app.NewButton()
	bu9.Text = "9"

	budive := app.NewButton()
	budive.Text = "/"
	
	bu0 := app.NewButton()
	bu0.Text = "0"
	
	buequile := app.NewButton()
	buequile.Text = "="
	
	buTimes := app.NewButton()
	buTimes.Text = "*"
	
	
	boradDefulte := app.NewBoard()
	app.DefulteBoard = boradDefulte
	
	Buttons(text,bu, bu2, bu3, buplus, bu4, bu5, bu6, buMinus, bu7,bu8,bu9,budive,bu0,buequile,buTimes)
	boradDefulte.AddChild(text, bu, bu2, bu3, buplus, bu4, bu5, bu6, buMinus,bu7,bu8,bu9,budive,bu0,buequile,buTimes)
	g.Update(app, func() {
		text.Width = app.GetScreenWidth() - 20
		
	})
}

func Buttons(input *g.Input,bu ...*g.Button) {
	x := 0
	xx := 0
	for index, button := range bu {
		// 1 2 3 +
		button.X = 5 + int32(index*73)
		button.Y = 100
		// 4 5 6 -
		if index == 4 || index == 5 || index == 6 || index == 7 {
			button.X = 5 + int32(x*73)
			button.Y = 155
			x++
		}
		// 7 8 9 /
		if index == 8 || index == 9 || index == 10 || index == 11 {
			button.X = 5 + int32(xx*73)
			button.Y = 210
			xx++
		}
		
		// 0
		if index == 12 {
			button.X = 5 
			button.Y = 265
		}
		// = 
		button.Width = 70

		if index == 13 {
			button.X = 80 
			button.Y = 265
			button.Width = 140
		}
		// *
		if index == 14 {
			button.X = 225 
			button.Y = 265
			button.Width = 70
		}
	}
}

func setNumber(number string){
	total = append(total,number)
}