package gox

import (
	"github.com/ncruces/zenity"
)

type Explorer struct{}

func (ex *Explorer) SelectFile(title string) (string, error) {
	return zenity.SelectFile(zenity.Title(title))
}

func (ex *Explorer) SelectFileMultiple(title string) ([]string, error) {
	return zenity.SelectFileMultiple(zenity.Title(title))
}

func (ex *Explorer) SelectDirctory(title string) (string, error) {
	return zenity.SelectFile(zenity.Title(title), zenity.Directory())
}

func (ex *Explorer) SelectDirctoryMultiple(title string) ([]string, error) {
	return zenity.SelectFileMultiple(zenity.Title(title), zenity.Directory())
}

func (ex *Explorer) SelectFileSave(title string, name string) (string, error) {
	return zenity.SelectFileSave(zenity.Title(title), zenity.ConfirmOverwrite(), zenity.Filename(name))
}
