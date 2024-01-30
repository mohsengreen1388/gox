package codepointtext

import (
	"github.com/benoitkugler/textlayout/fonts/glyphsnames"
	"github.com/go-text/typesetting/font"
	"github.com/go-text/typesetting/harfbuzz"
	"github.com/go-text/typesetting/language"
	Gobidi "github.com/lutzky/go-bidi"
	"golang.org/x/text/unicode/bidi"
	"os"
	"strconv"
	"strings"
)

type CodepointText struct {
	Text                    string
	Face                    harfbuzz.Face
	Font                    *harfbuzz.Font
	Buffer                  *harfbuzz.Buffer
	textsum                 string
	CodePoints              string
	Bidi                    Direction
	textAfterBidi           string
	codepointWithU          string
	codepointWithUni        rune
	convertToStringForPrint string
	err                     error
	displayer               Gobidi.Displayer
}

// setup up harfbuzz font,face,buffer
func (co *CodepointText) Init(fontFile string) {
	co.Buffer = harfbuzz.NewBuffer()
	fontload, err := os.Open(fontFile)
	defer fontload.Close()
	errorCodepintText(err, "font not load")
	co.Face, err = font.ParseTTF(fontload)
	errorCodepintText(err, "face has error")
	co.Font = harfbuzz.NewFont(co.Face)
	co.Buffer.Props.Script = language.Arabic
	co.Buffer.Props.Language = language.NewLanguage("ar")
	co.Buffer.Props.Direction = harfbuzz.LeftToRight
}

// convert text to codepoint
func (co *CodepointText) GetCodePointText() string {

	co.Buffer.Clear()
	co.Buffer.Info = nil
	co.textAfterBidi = ""
	co.codepointWithU = ""
	co.codepointWithUni = rune(0)
	co.convertToStringForPrint = ""
	co.textsum = ""
	co.textAfterBidi, _ = co.setbidi(co.Text, co.Bidi)

	co.Buffer.AddRunes([]rune(co.textAfterBidi), 0, -1)
	co.Buffer.Shape(co.Font, []harfbuzz.Feature{})

	for _, glyph := range co.Buffer.Info {
		co.codepointWithUni, _ = glyphsnames.GlyphToRune(co.Face.GlyphName(glyph.Glyph))
		co.codepointWithU = strings.Replace(string(co.codepointWithUni), "U+", "\\u", -1)
		co.convertToStringForPrint, co.err = strconv.Unquote("`"+co.codepointWithU+"`")
		errorCodepintText(co.err, "convertToStringForPrint is fail")
		co.makeCodePoints(co.convertToStringForPrint)
		co.textsum += co.convertToStringForPrint
	}
	return co.textsum
}

func errorCodepintText(err error, text string) {
	if err != nil {
		panic(text + err.Error())
	}
}

func (co *CodepointText) makeCodePoints(codepoint string) {
	co.CodePoints = ""
	co.CodePoints += "," + codepoint + ","
}

type Direction int

const (
	// LeftToRight indicates the text contains no right-to-left characters and
	// that either there are some left-to-right characters or the option
	// DefaultDirection(LeftToRight) was passed.
	LeftToRight Direction = iota

	// RightToLeft indicates the text contains no left-to-right characters and
	// that either there are some right-to-left characters or the option
	// DefaultDirection(RightToLeft) was passed.
	RightToLeft

	// Mixed indicates text contains both left-to-right and right-to-left
	// characters.
	Mixed

	// Neutral means that text contains no left-to-right and right-to-left
	// characters and that no default direction has been set.
	Neutral
)

func (co *CodepointText) setbidi(text string, bidiPos Direction) (string, error) {
	co.displayer = Gobidi.Displayer{BaseDir: bidi.Direction(bidiPos)}
	return co.displayer.Display(text)
}
