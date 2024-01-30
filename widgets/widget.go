package gox

import (
	"time"

	"github.com/ncruces/zenity"
)

type Widget struct{}

func (wi *Widget) Entry(text string) (string, error) {
	return zenity.Entry(text)
}

func (wi *Widget) Progress() (zenity.ProgressDialog, error) {
	return zenity.Progress()
}

func (wi *Widget) ListMultiple(text string, list []string) ([]string, error) {
	return zenity.ListMultiple(text, list)
}

func (wi *Widget) List(text string, list []string) (string, error) {
	return zenity.List(text, list)
}

func (wi *Widget) ListMultipleItems(text string) ([]string, error) {
	return zenity.ListMultipleItems(text)
}

func (wi *Widget) Password() (string, string, error) {
	return zenity.Password()
}

func (wi *Widget) Question(text string) error {
	return zenity.Question(text)
}

func (wi *Widget) Calendar(text string) (time.Time, error) {
	return zenity.Calendar(text)
}
