package cwidget

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
	"automateeverything.com/v2/template"
)

type testGroupWidget struct {
	ptg *template.TestGroup
}

func (c *testGroupWidget) NewTestGroupWidget(t *template.TestGroup) *testGroupWidget {
	return &testGroupWidget{ptg: t}
}

func (c * testGroupWidget) GetWidget() fyne.CanvasObject {
	return widget.NewLabel(c.ptg.TestGroupName)
}