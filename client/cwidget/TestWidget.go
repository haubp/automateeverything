package cwidget

import (
	"fyne.io/fyne/v2"
)

type TestWidget interface {
	GetWidget() fyne.CanvasObject
}