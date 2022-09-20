package template

import (
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
	"image/color"
)

// TestCategory test category
type TestCategory struct {
	TestCategoryName string `json:"test_category_name"`
	TestCategoryGroups []*TestGroup `json:"test_category_groups"`
	DisplayTestGroups bool `json:"-"`
	Widget *fyne.Container `json:"-"`
	W fyne.Window `json:"-"`
	A fyne.App `json:"-"`
}

// InitContext init test category
func (c * TestCategory) InitContext(a fyne.App, w fyne.Window) {
	c.DisplayTestGroups = false
	c.W = w
	c.A = a
	c.Widget = container.New(layout.NewHBoxLayout(),
		canvas.NewText(" Category:  " + c.TestCategoryName, color.RGBA{0xD8, 0xD8, 0xD8, 1}), 
		layout.NewSpacer(), 
		widget.NewButton("...", func(){
			c.DisplayTestGroups = !c.DisplayTestGroups
			if c.DisplayTestGroups {
				for i := range c.TestCategoryGroups {
					c.TestCategoryGroups[i].Widget.Show()
					c.TestCategoryGroups[i].DisplayTestCases = false
				}
			} else {
				for i := range c.TestCategoryGroups {
					c.TestCategoryGroups[i].Widget.Hide()
					c.TestCategoryGroups[i].HideAllChild()
				}
			}
		}),
		widget.NewButton("+", func(){
			newW := c.A.NewWindow("Add New Test Group")
			testGroupNameEntry := widget.NewEntry()
			content := container.New(	layout.NewGridLayoutWithRows(2),
										container.New(	layout.NewVBoxLayout(),
														testGroupNameEntry,
														layout.NewSpacer(),
										),
										container.New(	layout.NewVBoxLayout(),
														layout.NewSpacer(),
														widget.NewButton("Add", func(){
															newGroup := TestGroup{TestGroupName: testGroupNameEntry.Text, TestGroupTestCases: []*TestCase{} }
															newGroup.InitContext(c.A, c.W, c)
															newGroup.Widget.Show()
															c.TestCategoryGroups = append(c.TestCategoryGroups, &newGroup)

															UpdateUI(c.A, c.W, c)

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

// GetWidget get test category widget
func (c * TestCategory) GetWidget() *fyne.Container {
	return c.Widget
}

// DeleteTestGroup delete test group by index
func (c * TestCategory) DeleteTestGroup(i int) {
	c.TestCategoryGroups = append(c.TestCategoryGroups[:i], c.TestCategoryGroups[i+1:]...)
}