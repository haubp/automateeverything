package template

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"image/color"
)

// TestCase test case
type TestCase struct {
	TestCaseName string `json:"name"`
	TestCaseSteps []Step `json:"test_steps"`
	Widget *fyne.Container `json:"-"`
	A fyne.App `json:"-"`
	W fyne.Window `json:"-"`
	DisplayTestSteps bool `json:"-"`
}

// InitContext init test case
func (c * TestCase) InitContext(a fyne.App, w fyne.Window) {
	c.DisplayTestSteps = true
	c.W = w
	c.A = a
	c.Widget = container.New(layout.NewHBoxLayout(),
		container.NewMax(canvas.NewRectangle(color.RGBA{0, 100, 0, 1}), widget.NewLabel(c.TestCaseName)),
		layout.NewSpacer(),
		widget.NewButton("...", func(){
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
		widget.NewButton("+", func(){}),
	)
}

// HideAll test steps
func (c * TestCase) HideAll() {
	c.Widget.Hide()
	for i := range c.TestCaseSteps {
		if c.TestCaseSteps[i].Widget != nil {
			c.TestCaseSteps[i].Widget.Hide()
		}
	}
}

// ShowAll test steps
func (c * TestCase) ShowAll() {
	c.Widget.Show()
	for i := range c.TestCaseSteps {
		if c.TestCaseSteps[i].Widget != nil {
			c.TestCaseSteps[i].Widget.Show()
		}
	}
}

// GetWidget get test case widget
func (c * TestCase) GetWidget() *fyne.Container {
	return c.Widget
}

// DeleteTestStep delete test step by index
func (c * TestCase) DeleteTestStep(i int) {
	c.TestCaseSteps = append(c.TestCaseSteps[:i], c.TestCaseSteps[i+1:]...)
}