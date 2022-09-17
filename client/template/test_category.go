package template

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
)

type TestCategory struct {
	TestCategoryName string `json:"test_category_name"`
	TestCategoryGroups []TestGroup `json:"test_category_groups"`
	DisplayTestGroups bool
	Widget *fyne.Container
	W *fyne.Window
}

func (c * TestCategory) InitContext(w *fyne.Window) {
	c.DisplayTestGroups = true
	c.W = w
	c.Widget = container.New(layout.NewHBoxLayout(), 
		widget.NewButton(c.TestCategoryName, func(){
			c.DisplayTestGroups = !c.DisplayTestGroups
			if c.DisplayTestGroups {
				for i := range c.TestCategoryGroups {
					c.TestCategoryGroups[i].ShowAll()
				}
			} else {
				for i := range c.TestCategoryGroups {
					c.TestCategoryGroups[i].HideAll()
				}
			}
		}), 
		layout.NewSpacer(), 
		widget.NewButton("x", func(){}),
	)
}

func (c * TestCategory) GetWidget() *fyne.Container {
	return c.Widget
}

func (c * TestCategory) DeleteTestGroup(i int) {
	c.TestCategoryGroups = append(c.TestCategoryGroups[:i], c.TestCategoryGroups[i+1:]...)
}