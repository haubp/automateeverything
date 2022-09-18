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
	TestGroupTestCases []TestCase `json:"test_group_tests"`
	Category *TestCategory `json:"-"`
	Widget *fyne.Container `json:"-"`
	W fyne.Window `json:"-"`
	A fyne.App `json:"-"`
	DisplayTestCases bool `json:"-"`
}

// InitContext Init context for test group
func (c * TestGroup) InitContext(a fyne.App, w fyne.Window, t *TestCategory) {
	c.DisplayTestCases = true
	c.W = w
	c.A = a
	c.Category = t
	c.Widget = container.New(layout.NewHBoxLayout(),
		container.NewMax(canvas.NewRectangle(color.RGBA{0, 0, 100, 1}), widget.NewLabel(c.TestGroupName)),
		layout.NewSpacer(),
		widget.NewButton("...", func(){
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
															newTestCase := TestCase{TestCaseName: testCaseNameEntry.Text, TestCaseSteps: make([]Step, 0) }
															newTestCase.InitContext(c.A, c.W)
															c.TestGroupTestCases = append(c.TestGroupTestCases, newTestCase)

															// TODO: UPDATE HERE
															// Run Test Page
															tabs := container.NewAppTabs(
																container.NewTabItem("Create Test", CreateTestPage(c.A, c.Category)),
																container.NewTabItem("Run Test", CreateRunTestPage(c.A, c.Category)),
															)
															tabs.SetTabLocation(container.TabLocationLeading)
														
															c.W.CenterOnScreen()
															w.SetContent(tabs)

															newW.Close()
														}),
										),
									)
			newW.SetContent(content)
			newW.CenterOnScreen()
			newW.Resize(fyne.NewSize(300, 100))
			newW.Show()
		}),
	)
}

// HideAll hide all child test cases
func (c * TestGroup) HideAll() {
	c.Widget.Hide()
	for i := range c.TestGroupTestCases {
		if c.TestGroupTestCases[i].Widget != nil {
			c.TestGroupTestCases[i].Widget.Hide()
		}
		
	}
}

// ShowAll show all child test cases
func (c * TestGroup) ShowAll() {
	c.Widget.Show()
	for i := range c.TestGroupTestCases {
		if c.TestGroupTestCases[i].Widget != nil {
			c.TestGroupTestCases[i].Widget.Show()
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