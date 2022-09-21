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
	Result string `json:"result"`
	Widget *fyne.Container `json:"-"`
	A fyne.App `json:"-"`
	W fyne.Window `json:"-"`
	Category *TestCategory `json:"-"`
}

// InitContext init test case
func (c * TestCase) InitContext(a fyne.App, w fyne.Window, t *TestCategory) {
	c.W = w
	c.A = a
	c.Category = t
	c.Result = "Not Run"
	c.Widget = container.New(layout.NewHBoxLayout(),
		canvas.NewText(" \t\t\t\t" + c.TestCaseName, color.RGBA{0x64, 0x7E, 0x68, 1}),
		layout.NewSpacer(),
		widget.NewButton("Select", func(){
			SelectedTestCase = c
			UpdateUI(c.A, c.W, c.Category)
		}),
	)
	c.Widget.Hide()
}

// GetWidget get test case widget
func (c * TestCase) GetWidget() *fyne.Container {
	return c.Widget
}

// DeleteTestStep delete test step by index
func (c * TestCase) DeleteTestStep(i int) {
	c.TestCaseSteps = append(c.TestCaseSteps[:i], c.TestCaseSteps[i+1:]...)
}