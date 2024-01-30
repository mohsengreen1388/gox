package main

import (
	gox "github.com/mohsengreen1388/gox/widgets"
)

func main() {

	app := gox.Init("App", 700, 400)

		button := app.NewButton()
		button.Text = "ok"
		button.BackgroundColor = gox.BLUE
		button2 := app.NewButton()
		button2.Text = "cancel"
		button2.BackgroundColor = gox.PINK
		button2.Icon = gox.AddVideo
		button2.FuncExce = func() {
			// your code run
		}

	gox.Update(app, func() {
		button.Draw(10,10,200,50)
		button2.Draw(10,80,200,50)
	})
}
