package template

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2"
)

type TestGroup struct {
	TestGroupName string `json:"test_group_name"`
	TestGroupTestCases []TestCase `json:"test_group_tests"`
	Widget *fyne.Container
	W *fyne.Window
	DisplayTestCases bool
}

func (c * TestGroup) InitContext(w *fyne.Window) {
	c.DisplayTestCases = true
	c.W = w
	c.Widget = container.New(layout.NewHBoxLayout(),
		widget.NewButton("\t\t- " + c.TestGroupName, func(){
			c.DisplayTestCases = !c.DisplayTestCases
			if c.DisplayTestCases {
				for i := range c.TestGroupTestCases {
					c.TestGroupTestCases[i].ShowAll()
				}
			} else {
				for i := range c.TestGroupTestCases {
					c.TestGroupTestCases[i].HideAll()
				}
			}
		}),
		layout.NewSpacer(),
		widget.NewButton("x", func(){}),
	)
}

func (c * TestGroup) HideAll() {
	c.Widget.Hide()
	for i := range c.TestGroupTestCases {
		c.TestGroupTestCases[i].Widget.Hide()
	}
}

func (c * TestGroup) ShowAll() {
	c.Widget.Show()
	for i := range c.TestGroupTestCases {
		c.TestGroupTestCases[i].Widget.Show()
	}
}

func (c * TestGroup) GetWidget() *fyne.Container {
	return c.Widget
}

func (c * TestGroup) DeleteTestCase(i int) {
	c.TestGroupTestCases = append(c.TestGroupTestCases[:i], c.TestGroupTestCases[i+1:]...)
}