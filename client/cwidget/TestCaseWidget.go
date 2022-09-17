package cwidget

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
	"automateeverything.com/v2/template"
)

type testCaseWidget struct {
	pt *template.TestCase
}

func (c *testCaseWidget) NewCategoryWidget(t *template.TestCase) *testCaseWidget {
	return &testCaseWidget{pt: t}
}

func (c * testCaseWidget) GetWidget() fyne.CanvasObject {
	return widget.NewLabel(c.pt.TestName)
}