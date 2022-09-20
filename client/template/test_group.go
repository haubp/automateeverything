package template

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"image/color"
)

// TestGroup test group
type TestGroup struct {
	TestGroupName string `json:"test_group_name"`
	TestGroupTestCases []*TestCase `json:"test_group_tests"`
	Category *TestCategory `json:"-"`
	Widget *fyne.Container `json:"-"`
	W fyne.Window `json:"-"`
	A fyne.App `json:"-"`
	DisplayTestCases bool `json:"-"`
}

// InitContext Init context for test group
func (c * TestGroup) InitContext(a fyne.App, w fyne.Window, t *TestCategory) {
	c.DisplayTestCases = false
	c.W = w
	c.A = a
	c.Category = t
	c.Widget = container.New(layout.NewHBoxLayout(),
		canvas.NewText(" \t\t" + c.TestGroupName, color.RGBA{0x56, 0x2B, 0x08, 1}),
		layout.NewSpacer(),
		widget.NewButton("...", func(){
			c.DisplayTestCases = !c.DisplayTestCases
			if c.DisplayTestCases {
				for i := range c.TestGroupTestCases {
					c.TestGroupTestCases[i].Widget.Show()
				}
			} else {
				for i := range c.TestGroupTestCases {
					c.TestGroupTestCases[i].Widget.Hide()
				}
			}
		}),
		widget.NewButton("+", func(){
			newW := c.A.NewWindow("Add New Test Case")
			testCaseNameEntry := widget.NewEntry()
			content := container.New(	layout.NewGridLayoutWithRows(2),
										container.New(	layout.NewVBoxLayout(),
										testCaseNameEntry,
														layout.NewSpacer(),
										),
										container.New(	layout.NewVBoxLayout(),
														layout.NewSpacer(),
														widget.NewButton("Add", func(){
															newTestCase := TestCase{TestCaseName: testCaseNameEntry.Text, TestCaseSteps:[]Step{} }
															newTestCase.InitContext(c.A, c.W, c.Category)
															newTestCase.Widget.Show()
															c.TestGroupTestCases = append(c.TestGroupTestCases, &newTestCase)

															UpdateUI(c.A, c.W, c.Category)

															newW.Close()
														}),
										),
									)
			newW.SetContent(content)
			newW.CenterOnScreen()
			newW.Resize(fyne.NewSize(300, 100))
			newW.Show()
		}),
		widget.NewButton("X", func() {
			
		}),
	)
	c.Widget.Hide()
}

// HideAllChild hide all child test cases
func (c * TestGroup) HideAllChild() {
	c.Widget.Hide()
	for i := range c.TestGroupTestCases {
		if c.TestGroupTestCases[i].Widget != nil {
			c.TestGroupTestCases[i].Widget.Hide()
		}
	}
}

// GetWidget get test group widget
func (c * TestGroup) GetWidget() *fyne.Container {
	return c.Widget
}

// DeleteTestCase delete test case by index
func (c * TestGroup) DeleteTestCase(i int) {
	c.TestGroupTestCases = append(c.TestGroupTestCases[:i], c.TestGroupTestCases[i+1:]...)
}