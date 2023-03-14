package template

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// UpdateUI update test page UI
func UpdateUI(a fyne.App, w fyne.Window, t *TestCategory) {
	tabs := container.NewAppTabs(
		container.NewTabItem("Create Test", CreateTestPage(a, w, t)),
		container.NewTabItem("Run Test", RunTestPage(a, w)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
}