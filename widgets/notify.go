package gox

import (
	"github.com/ncruces/zenity"
)

type Notify struct{}

func (no *Notify) Notify(text string) error {
	return zenity.Notify(text)
}

func (no *Notify) Warning(text string) error {
	return zenity.Warning(text)
}

func (no *Notify) Info(text string) error {
	return zenity.Info(text)
}

func (no *Notify) Error(text string) error {
	return zenity.Error(text)
}
