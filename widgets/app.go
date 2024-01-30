package gox

import (
	"encoding/base64"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/mohsengreen1388/gox/pkg/codepointtext"
	"image/color"
)

type App struct {
	Font                   rl.Font
	FontAddreesFlie        string
	FontSize               int32
	SlideSize, MenuBarSize int32
	RadioButtonChecker     map[string]*RadioButton
	BackgroundColor        rl.Color
	DefulteBoard           *Board
	Lock                   bool
	Icon                   *Icon
	CodePoint              *CodePoint
}

// we can use by this interface pass draw to other object
type InterfaceDraw interface {
	Coord() rl.Rectangle
	DrawInterface(x, y, width, hight int32)
	EnableMemberModal()
}

// App is main object that we can by it call other object
// examle app.NewButton()
func newApp() *App {
	return &App{
		Font:               rl.GetFontDefault(),
		BackgroundColor:    rl.White,
		RadioButtonChecker: make(map[string]*RadioButton),
		Icon:               NewIcon(),
	}
}

func (a *App) NewTopBar(logoName, tabName string) *TopBar {
	button := a.NewButton()
	button.Icon = Menu
	button.IconSize = 50
	button.iconVector = rl.Vector2{-27, -17}
	button.IconColor = rl.SkyBlue

	return &TopBar{
		Font:          a.Font,
		app:           a,
		Text:          tabName,
		FontSize:      34,
		TextColor:     rl.Black,
		App:           a,
		Button:        button,
		TextLogo:      logoName,
		TextLogoSize:  24,
		TextLogoColor: rl.SkyBlue,
		ImageWidth:    float32(100),
		Imagehight:    float32(100),
		IconTop:       Home,
		iconSize:      64,
		IconColor:     rl.SkyBlue,
	}
}

// New text input
func (a *App) NewInput() *Input {
	return &Input{
		Font:                a.Font,
		app:                 a,
		FontSize:            24,
		borderColor:         rl.Black,
		BorderColorActive:   rl.Blue,
		BorderColorUnactive: rl.Black,
		TextColor:           rl.Black,
		PlaceholderDefulte:  "typeing...",
		Width:               200,
		Hight:               50,
		space:               2,
	}
}

// New text button
func (a *App) NewButton() *Button {
	return &Button{
		Width:            200,
		Hight:            50,
		textPostionStart: 0,
		BackgroundColor:  rl.Blue,
		TextColor:        rl.RayWhite,
		FontSize:         24,
		Text:             "Button",
		Font:             a.Font,
		round:            0.3,
		app:              a,
		IconSize:         23,
		IconColor:        rl.White,
		iconVector:       rl.Vector2{X: -32, Y: -4},
	}
}

// Init icon
func NewIcon() *Icon {
	iconData := IconData{}
	data := iconData.Data()
	dataDecode, _ := base64.StdEncoding.DecodeString(string(data))
	image := rl.LoadImageFromMemory(".png", dataDecode, int32(len(dataDecode)))
	texture := rl.LoadTextureFromImage(image)
	defer unloadIcon(&dataDecode, image)

	return &Icon{
		dataFileIcon:   data,
		iconDataDecode: dataDecode,
		image:          image,
		texture:        texture,
	}
}

// New image
func (a *App) NewImage(file string, width, hight int) *Image {
	image := rl.LoadImage(file)
	rl.ImageResize(image, int32(width), int32(hight))
	defer rl.UnloadImage(image)

	return &Image{
		Texture: rl.LoadTextureFromImage(image),
		app:     a,
		Width:   int32(width),
		Hight:   int32(hight),
	}
}

// New label
func (a *App) Newlabel() *Label {
	return &Label{
		Font:       a.Font,
		TextColor:  rl.Black,
		TextSize:   24,
		Textspace:  2,
		app:        a,
		IconColor:  rl.Black,
		iconVector: rl.Vector2{-4, 0},
	}
}

// New modal
func (a *App) NewModal() *Modal {
	return &Modal{
		Font:            a.Font,
		app:             a,
		Rounde:          0.1,
		BackgroundColor: rl.White,
		shadowColor:     rl.Gray,
	}
}

// New SidePanel be side right and left
func (a *App) NewSidePanel() *SidePanel {
	model := a.NewModal()
	return &SidePanel{
		Font:            a.Font,
		app:             a,
		model:           model,
		Rounde:          0.1 / 2,
		BackgroundColor: rl.White,
		shadowColor:     rl.Gray,
		Width:           250,
		Hight:           int32(rl.GetMonitorHeight(rl.GetCurrentMonitor())) - 200,
		Y:               100,
		Direction:       Right,
	}
}

// we can this object multi board
// per board like a canvas
func (a *App) NewBoard() *Board {
	win := &Board{
		Font: a.Font,
		app:  a,
	}
	if win.app.DefulteBoard == nil {
		a.DefulteBoard = win
	}
	return win
}

// New tab button
func (a *App) NewTab(count int8, childBoard bool) *TabBar {
	tab := &TabBar{
		Font:            a.Font,
		app:             a,
		ChildBoard:      childBoard,
		Count:           int8(count),
		AutoSize:        true,
		Minsize:         75,
		SpeedScroll:     2,
		X:               0,
		Y:               50,
		Width:           200,
		Hight:           40,
		FontSize:        20,
		TextColor:       rl.Black,
		BackgroundColor: rl.Blue,
	}
	tab.AddChildButtons()
	return tab
}

// New Progress
func (a *App) NewProgress(value, width int32) *Progress {
	return &Progress{
		Font:           a.Font,
		app:            a,
		Color:          rl.SkyBlue,
		Width:          width,
		Hight:          25,
		Value:          value,
		ShowPrecentage: true,
	}
}

func (a *App) NewCheckBox() *CheckBox {
	black := rl.Black
	black.A = 80
	return &CheckBox{
		Font:            a.Font,
		app:             a,
		BackgroundColor: black,
		CircleColor:     rl.White,
		Hight:           25,
		Width:           54,
	}
}

func (a *App) NewRadioButton(name string, value interface{}, id string) *RadioButton {
	black := rl.Black
	black.A = 80
	lable := a.Newlabel()
	lable.Text = name

	RadioButton := &RadioButton{
		Font:               a.Font,
		app:                a,
		BackgroundColor:    black,
		CircleColor:        black,
		CircleColorChecked: rl.Green,
		Hight:              25,
		Width:              54,
		radius:             13,
		label:              *lable,
		Value:              value,
		Name:               name,
		Id:                 id,
		TextColor:          rl.Black,
	}

	a.RadioButtonChecker[RadioButton.Name+RadioButton.Id] = RadioButton

	return RadioButton
}

func (a *App) NewSliderBar(min, max, width int32) *SlideBar {
	return &SlideBar{
		Font:                a.Font,
		app:                 a,
		Color:               rl.SkyBlue,
		Width:               width,
		Hight:               25,
		Min:                 min,
		Max:                 max,
		Value:               min,
		CircleRangeColor:    rl.Orange,
		PercentageColor:     rl.SkyBlue,
		TextPercentageColor: rl.Black,
		textPercentageSize:  16,
	}
}

func (a *App) NewMeneFloat() *MenuFloat {
	model := a.NewModal()
	label := a.Newlabel()
	labelWithoutIcon := a.Newlabel()
	labelMainIteamIcon := a.Newlabel()
	labelShortKey := a.Newlabel()
	labelMainIteamIcon.Icon = ArrowRight

	return &MenuFloat{
		BackgroundColor:    model.BackgroundColor,
		LineColor:          rl.Black,
		model:              model,
		label:              label,
		labelWithoutIcon:   labelWithoutIcon,
		labelMainIteamIcon: labelMainIteamIcon,
		labelShortKey:      labelShortKey,
		Width:              250,
		Hight:              30,
		HightPerIteam:      30,
		offset:             25,
		Enable:             true,
	}
}

func (a *App) NewMeneBar() *MenuBar {
	BackgroundColor := color.RGBA{120, 120, 120, 120}
	return &MenuBar{
		BackgroundColor:     BackgroundColor,
		Font:                a.Font,
		FontSize:            20,
		app:                 a,
		x:                   0,
		y:                   0,
		hight:               20,
		TextColor:           rl.Black,
		spaceBetwenPerIteam: 10,
	}
}

func (a *App) drawDefultBoard() {
	a.DefulteBoard.Draw(0, 0, a.GetScreenWidth(), a.GetScreenHeight())
}

func (a *App) NewExplorer() *Explorer {
	return &Explorer{}
}

func (a *App) NewWidget() *Widget {
	return &Widget{}
}
func (a *App) NewColorPicker() *ColorPicker {
	return &ColorPicker{}
}

func (a *App) NewNotfiy() *Notify {
	return &Notify{}
}

// Init window if you want use wihout raylib-go
func Init(title string, width, hight int32) *App {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(width, hight, title)

	return newApp()
}

// Exce to update a for per frame
func Update(a *App, f func()) {
	rl.SetTargetFPS(25)
	for !rl.WindowShouldClose() {
		rl.ClearBackground(a.BackgroundColor)
		rl.BeginDrawing()
		f()
		if a.DefulteBoard != nil {
			a.drawDefultBoard()
		}
		rl.EndDrawing()
	}
	a.Unload()
	rl.CloseWindow()
}

func (a *App) LoadFont(file string, size int32) rl.Font {
	a.FontAddreesFlie = file
	a.FontSize = size
	a.CodePoint = NewCodePoints(file)
	a.CodePoint.Init()
	return rl.LoadFontEx(file, size, nil)
}

func NewCodePoints(file string) *CodePoint {
	return &CodePoint{
		codepoint: &codepointtext.CodepointText{},
		file:      file,
	}
}

func (a *App) CheckForNotMemberSideBar(object bool) int32 {
	if object {
		return 0
	}
	return a.SlideSize
}

func (a *App) GetScreenWidth() int32 {
	return int32(rl.GetScreenWidth() - int(a.SlideSize))
}

func (a *App) GetScreenHeight() int32 {
	return int32(rl.GetScreenHeight())
}

func (a *App) MouseX() int32 {
	return rl.GetMouseX()
}

func (a *App) MouseY() int32 {
	return rl.GetMouseY()
}

func (a *App) GetScreenHeightByprcentage(value float32) float32 {
	prcent := value * float32(a.GetScreenHeight()) / float32(100)
	return prcent
}

func (a *App) GetScreenWidthByprcentage(value float32) float32 {
	prcent := value * float32(a.GetScreenWidth()) / float32(100)
	return prcent
}

func unloadIcon(dataicon *[]byte, im *rl.Image) {
	dataicon = nil
	rl.UnloadImage(im)
}

func (a *App) Unload() {
	a.Icon.Unload()
	rl.UnloadFont(a.Font)
}
