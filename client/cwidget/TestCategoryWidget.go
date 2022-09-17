package cwidget

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2"
	"automateeverything.com/v2/template"
)

type testCategoryWidget struct {
	pt *template.TestCategory
}

func (c *testCategoryWidget) NewCategoryWidget(t *template.TestCategory) *testCategoryWidget {
	return &testCategoryWidget{pt: t}
}

func (c * testCategoryWidget) GetWidget() fyne.CanvasObject {
	return widget.NewLabel(c.pt.TestCategoryName)
}