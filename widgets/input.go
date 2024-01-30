package gox

import (
	"github.com/mohsengreen1388/gox/pkg/codepointtext"
	rl "github.com/gen2brain/raylib-go/raylib"
	"golang.org/x/exp/slices"
	"strings"
)

type Input struct {
	X, Y, Width, Hight     int32
	RectangleBody          rl.Rectangle
	checkIsActive          bool
	text                   string
	textSlice              []string
	mouseX                 int
	mouseY                 int
	pointermoveForPressKey float32
	textPostion            rl.Vector2
	textDirectionSub       float32
	mesureCounter          []float32
	lenText                int
	LimitText              int32
	Placeholder            string
	PlaceholderDefulte     string
	PlaceholderColor       rl.Color
	selectRecPostion       rl.Vector2
	selectRecSize          rl.Vector2
	textWidth              int
	space                  int8
	selectState            bool
	countWalke             int32
	Font                   rl.Font
	FontSize               float32
	Background             rl.Color
	TextColor              rl.Color
	borderColor            rl.Color
	BorderColorActive      rl.Color
	BorderColorUnactive    rl.Color
	app                    *App
	ReadOnly               bool
	MemberModal            bool
}

func (in *Input) Draw(x, y, width, hight int32) {
	in.mouseX = int(rl.GetMouseX())
	in.mouseY = int(rl.GetMouseY())
	in.PlaceholderColor.A = 120
	in.body(x, y, width, hight)
}

func (in *Input) body(x, y, width, hight int32) {
	recX := float32(x)
	recY := float32(y)
	recWidth := float32(width)
	recHight := float32(hight)
	space := float32(15)
	in.textPostion = rl.Vector2{float32(recX + float32(space) - in.textDirectionSub), float32(recY + space)}
	if recX == 0 {
		recX = 5
	} // margin right

	in.RectangleBody = rl.Rectangle{recX, recY, recWidth, recHight}
	rl.DrawRectangleRounded(in.RectangleBody, 0.3, 0, in.Background)
	rl.DrawRectangleRoundedLines(in.RectangleBody, 0.3, 0, 3, in.borderColor)
	rl.BeginScissorMode(in.RectangleBody.ToInt32().X, in.RectangleBody.ToInt32().Y, in.RectangleBody.ToInt32().Width, in.RectangleBody.ToInt32().Height)

	rl.DrawTextEx(in.Font, in.text, in.textPostion, in.FontSize, float32(in.space), in.TextColor)
	rl.DrawTextEx(in.Font, in.Placeholder, in.textPostion, in.FontSize, float32(in.space), in.PlaceholderColor)

	if in.checkIsActive {

		in.textSetInput(recX, recY, recWidth, recHight, space)
		in.pointerMove()
		in.deleteKey()
		// Todo list
		//in.SelectText()
	}
	rl.EndScissorMode()
	if !in.ReadOnly {
		in.event()
	}
}

func (in *Input) event() {
	if len(in.textSlice) == 0 {
		in.Placeholder = in.PlaceholderDefulte
	} else {
		in.Placeholder = ""
	}
	CheckCollisionRecsByMouseIsTrue := rl.CheckCollisionRecs(in.RectangleBody, rl.Rectangle{float32(in.mouseX), float32(in.mouseY), 0, 0})
	if CheckCollisionRecsByMouseIsTrue && in.isMemberModal() {
		in.movePointerByMouse()
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			in.checkIsActive = true
			in.borderColor = in.BorderColorActive
		}
	}

	if !CheckCollisionRecsByMouseIsTrue {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			in.checkIsActive = false
			in.borderColor = in.BorderColorUnactive
			rl.SetMouseCursor(rl.MouseCursorArrow)
		}
	}
}

func (in *Input) textSetInput(recX, recY, recWidth, recHight, space float32) {
	rl.SetMouseCursor(rl.MouseCursorIBeam)
	key := rl.GetCharPressed()
	if key > 20 && in.checkIsActive {
		if in.lenText < 0 {
			in.textSlice = append(in.textSlice, string(key))
		} else {
			in.textSlice = slices.Insert(in.textSlice, in.lenText, string(key))
		}
		measureFontKey := rl.MeasureTextEx(in.Font, string(key), in.FontSize, float32(in.space))
		in.pointermoveForPressKey += measureFontKey.X + float32(in.space)
		in.mesureCounter = slices.Insert(in.mesureCounter, in.lenText, measureFontKey.X+float32(in.space))
		in.lenText++
		in.text = strings.Join(in.textSlice, "")
		// Todo list
		// in.textencode()

		if int(in.pointermoveForPressKey) > int(in.RectangleBody.Width/1.1) {
			in.textDirectionSub -= -(measureFontKey.X + float32(in.space))
		}
	}
	rl.DrawLineEx(rl.Vector2{in.textPostion.X + float32(in.pointermoveForPressKey), recY + 12}, rl.Vector2{in.textPostion.X + float32(in.pointermoveForPressKey), recY + 40}, float32(in.space), rl.Black)
}

func (in *Input) pointerMove() {
	keyForMove := rl.GetKeyPressed()
	if keyForMove == rl.KeyRight || keyForMove == rl.KeyLeft {
		if keyForMove == rl.KeyRight {
			if in.lenText < len(in.textSlice) {
				in.pointermoveForPressKey += in.mesureCounter[in.lenText]
				if in.RectangleBody.Width <= in.pointermoveForPressKey {
					in.textDirectionSub += in.mesureCounter[in.lenText]
				}
				in.lenText++
			}
		}
		if keyForMove == rl.KeyLeft {
			if in.lenText > 0 {
				in.lenText--
				in.pointermoveForPressKey -= in.mesureCounter[in.lenText]
				if in.textDirectionSub > 0 {
					in.textDirectionSub -= in.mesureCounter[in.lenText]
				}
			}
		}
	}
}

// Todo list
func (in *Input) flash() {}

func (in *Input) movePointerByMouse() {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if !in.selectState {
			in.selectState = true
		} else {
			in.selectState = false
		}

		in.textWidth = 0
		for _, value := range in.mesureCounter {
			in.textWidth += int(value)
		}
		if in.mesureCounter != nil && len(in.textSlice) != 0 {
			coordinatesRec := int(in.mouseX) - int(in.textPostion.X)
			postionPerWord := coordinatesRec / (in.textWidth / len(in.textSlice))

			if postionPerWord > len(in.textSlice) {
				postionPerWord = len(in.textSlice)
			}
			if postionPerWord >= 0 {
				in.lenText = postionPerWord
				wholeTextThatSelectWidth := 0.0

				for _, vulue := range in.mesureCounter[0:postionPerWord] {
					wholeTextThatSelectWidth += float64(vulue)
				}

				in.pointermoveForPressKey = float32(wholeTextThatSelectWidth)
			}
		}
	}
}

func (in *Input) deleteKey() {
	if rl.IsKeyPressed(rl.KeyBackspace) {
		if in.lenText > 0 {
			in.lenText--
			if len(in.mesureCounter) >= 0 {
				in.pointermoveForPressKey -= in.mesureCounter[in.lenText]
			}
			sliceTextRemove := slices.Delete(in.textSlice, in.lenText, in.lenText+int(1))
			in.textSlice = sliceTextRemove
			sliceMesurRemove := slices.Delete(in.mesureCounter, in.lenText, in.lenText+int(1))
			in.mesureCounter = sliceMesurRemove

			in.text = strings.Join(in.textSlice, "")
		}
	}
}

// Todo list
func (in *Input) selectText() {
	in.selectRecPostion.Y = 22
	in.selectRecSize.Y = 28
	in.selectRecPostion.X = float32(in.pointermoveForPressKey + in.textPostion.X)
	mouseDirection := rl.GetMouseDelta()

	rl.DrawRectangleV(in.selectRecPostion, rl.Vector2{in.selectRecSize.X, in.selectRecSize.Y}, in.PlaceholderColor)

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		if in.selectState {
			in.selectRecPostion.X = 0
			in.selectRecSize.X = 0
			return
		}

		if (in.textPostion.X + float32(in.textWidth)) >= float32(in.mouseX) {
			if mouseDirection.X >= 0 && in.selectRecSize.X <= float32(in.textWidth) {
				if in.lenText <= len(in.mesureCounter) {
					in.selectRecSize.X = in.mesureCounter[in.lenText]
					in.selectRecPostion.X += 1
				}
			}
		}
		if mouseDirection.X <= 0 {
		}
	}
}

func (in *Input) SetPosition(x, y int32) {
	in.X = x
	in.Y = y
}

func (in *Input) Coord() rl.Rectangle {
	return rl.NewRectangle(float32(in.X), float32(in.Y), float32(in.Width), float32(in.Hight))
}

func (in *Input) DrawInterface(x, y, width, hight int32) {
	in.mouseX = int(rl.GetMouseX())
	in.mouseY = int(rl.GetMouseY())
	in.PlaceholderColor.A = 120
	in.body(x, y, width, hight)
}

func (in *Input) isMemberModal() bool {
	if in.MemberModal {
		return true
	}
	return !in.app.Lock
}

func (in *Input) EnableMemberModal() {
	in.MemberModal = true
}

func (in *Input) DisableMemberModal() {
	in.MemberModal = false
}

func (in *Input) SetText(text string) {
	in.clearText()
	for _, charKey := range text {
		if in.lenText < 0 {
			in.textSlice = append(in.textSlice, string(charKey))
		} else {
			in.textSlice = slices.Insert(in.textSlice, in.lenText, string(charKey))
		}
		measureFontKey := rl.MeasureTextEx(in.Font, string(charKey), in.FontSize, float32(in.space))
		in.pointermoveForPressKey += measureFontKey.X + float32(in.space)
		in.mesureCounter = slices.Insert(in.mesureCounter, in.lenText, measureFontKey.X+float32(in.space))
		in.lenText++
		in.text = strings.Join(in.textSlice, "")
	}
}

func (in *Input) clearText() {
	in.textSlice = nil
	in.text = ""
	in.lenText = 0
}

func (in *Input) GetText() string {
	return in.text
}

// Todo list
func (in *Input) textencode() {
	in.text = in.app.CodePoint.Genrate(in.text, codepointtext.LeftToRight)
}
