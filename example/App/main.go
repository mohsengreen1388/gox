package main

import (
	gox "github.com/mohsengreen1388/gox/widgets"
)

func main() {

	app := gox.Init("App",700,400)
		text := app.Newlabel()
		text.TextColor = gox.BLUE
		text.TextSize = 25	

	gox.Update(app,func() {
		text.Draw("Welcome to GoX",float32(app.GetScreenWidth())/2.5,float32(app.GetScreenHeight())/2.5)
	})
}