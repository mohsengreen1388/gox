package gox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Image struct {
	X, Y, Width, Hight int32
	Font               rl.Font
	Texture            rl.Texture2D
	app                *App
	MemberModal        bool
}

func (im *Image) Draw(x, y int32) {
	im.body(x, y)
}

func (im *Image) body(x, y int32) {
	rl.DrawTexture(im.Texture, x, y, rl.RayWhite)
}

func (im *Image) Unload() {
	rl.UnloadTexture(im.Texture)
}

func (im *Image) SetPosition(x, y, width int32) {
	im.X = x
	im.Y = y
}

func (im *Image) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(im.X), float32(im.Y), float32(im.Width), float32(im.Hight))
}

func (im *Image) DrawInterface(x, y, width, hight int32) {
	im.body(x, y)
}

func (im *Image) EnableMemberModal() {
	im.MemberModal = true
}
