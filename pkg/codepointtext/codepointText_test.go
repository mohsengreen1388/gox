package codepointtext

import (
	"testing"
)


func BenchmarkInit(b *testing.B){
	c := CodepointText{}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Init("./Vazir.ttf")
	}
	b.StopTimer()
}

func TestGetCodePointText(t *testing.T) {
	c := codepoint()
	c.Text = "abc"
	text := c.GetCodePointText()
	if text == "" {
		t.Fatalf("GetCodePointText not get codepoint")
	}
}

func BenchmarkGetCodePointText(b *testing.B){
	c := codepoint()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		c.Text = "سلام Hello"
		c.GetCodePointText()
	}
	b.StopTimer()
}

func codepoint()CodepointText{
	c := CodepointText{}
	c.Init("./Vazir.ttf")
	return c
}
