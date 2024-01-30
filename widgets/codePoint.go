package gox

import "github.com/mohsengreen1388/gox/pkg/codepointtext"

type CodePoint struct {
	codepoint *codepointtext.CodepointText
	file      string
}

func (c *CodePoint) Init() {
	c.codepoint.Init(c.file)
}

func (c *CodePoint) Genrate(text string, bidi codepointtext.Direction) string {
	c.codepoint.Text = text
	c.codepoint.Bidi = bidi
	return c.codepoint.GetCodePointText()
}
