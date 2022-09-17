package template

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
)

type TestCase struct {
	TestCaseName string `json:"name"`
	TestCaseSteps []Step `json:"test_steps"`
	Widget *fyne.Container
	W *fyne.Window
	DisplayTestSteps bool
}

func (c * TestCase) InitContext(w *fyne.Window) {
	c.DisplayTestSteps = true
	c.W = w
	c.Widget = container.New(layout.NewHBoxLayout(),
		widget.NewButton("\t\t\t\t* " + c.TestCaseName, func(){
			c.DisplayTestSteps = !c.DisplayTestSteps
			if c.DisplayTestSteps {
				for i := range c.TestCaseSteps {
					c.TestCaseSteps[i].Widget.Show()
				}
			} else {
				for i := range c.TestCaseSteps {
					c.TestCaseSteps[i].Widget.Hide()
				}
			}
		}),
		layout.NewSpacer(),
		widget.NewButton("x", func(){}),
	)
}

func (c * TestCase) HideAll() {
	c.Widget.Hide()
	for i := range c.TestCaseSteps {
		c.TestCaseSteps[i].Widget.Hide()
	}
}

func (c * TestCase) ShowAll() {
	c.Widget.Show()
	for i := range c.TestCaseSteps {
		c.TestCaseSteps[i].Widget.Show()
	}
}

func (c * TestCase) GetWidget() *fyne.Container {
	return c.Widget
}

func (c * TestCase) DeleteTestStep(i int) {
	c.TestCaseSteps = append(c.TestCaseSteps[:i], c.TestCaseSteps[i+1:]...)
}