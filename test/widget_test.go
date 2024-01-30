package test

import (
	"github.com/mohsengreen1388/gox/widgets"
	"testing"
)

var app *gox.App
var input *gox.Input
var board *gox.Board
var menuFloat *gox.MenuFloat

func init() {
	app = gox.Init("Test", 200, 200)
	input = app.NewInput()
	menuFloat = app.NewMeneFloat()
	board = app.NewBoard()
	app.DefulteBoard = board

}

// Test Board
func TestBoard(t *testing.T) {
	app.DefulteBoard.AddChild(input)
}

func BenchmarkBoard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		board := app.NewBoard()
		app.DefulteBoard = board
		app.DefulteBoard.AddChild(input)
	}
}

// Button test
func TestButton(t *testing.T) {
	bu := app.NewButton()
	bu.Draw(0, 0, 100, 100)
}

func BenchmarkButton(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bu := app.NewButton()
		bu.Draw(0, 0, 10, 10)
	}
}

// Input test
func TestInput(t *testing.T) {
	in := app.NewInput()
	in.Draw(0, 0, 100, 100)
}

func BenchmarkInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := app.NewInput()
		in.Draw(0, 0, 10, 10)
	}
}

func TestSetText(t *testing.T) {
	input.SetText("ok")
}

func BenchmarkSetText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input.SetText("ok")
	}
}

func TestGetText(t *testing.T) {
	input.GetText()
}

func BenchmarkGetText(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input.GetText()
	}
}

// Test CheckBox
func TestCheckBox(t *testing.T) {
	ch := app.NewCheckBox()
	ch.Draw(0, 0)
}

func BenchmarkCheckBox(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := app.NewCheckBox()
		ch.Draw(0, 0)
	}
}

// Test Image
func TestImage(t *testing.T) {
	im := app.NewImage("", 200, 200)
	im.Draw(0, 0)
}

func BenchmarkImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		im := app.NewImage("", 200, 200)
		im.Draw(0, 0)
	}
}

// Test Menu
func TestMenu(t *testing.T) {
	menuFloat.AddMainMenu("testMainIteam", gox.NoIcon, menuFloat)
	menuFloat.AddSubMenu("testSubIteam", gox.ArrowRight, "Ctrl+S", func() {})
	menu := app.NewMeneBar()
	menu.AddItem("Test", menuFloat)
	menu.Draw()
}

func BenchmarkMenu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		menuFloat.AddMainMenu("testMainIteam", gox.NoIcon, menuFloat)
		menuFloat.AddSubMenu("testSubIteam", gox.ArrowRight, "Ctrl+S", func() {})
		menu := app.NewMeneBar()
		menu.AddItem("Test", menuFloat)
		menu.Draw()
	}
}

