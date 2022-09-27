package template

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"

	"automateeverything.com/v2/custom"
)

// RunTestPage Creat Run Test Page widget
func RunTestPage(a fyne.App, w fyne.Window) *fyne.Container {
	var t TestCategory
	resultTelevision := custom.NewFixSizeLabel()
	resultTelevision.SetFixSize(fyne.NewSize(500, 100))
	testRunProgress := widget.NewProgressBar()
	testRunProgress.SetValue(0.0)
	reportNameEntry := widget.NewEntry()
	reportNameEntry.SetText("report.json")
	templateFileSelectedLabelIndicator := widget.NewLabel("No test template found")
	scrollResultContainer := container.NewVScroll(
		resultTelevision,
	)
	runTestButton := widget.NewButton("Run", func() {
		if len(t.TestCategoryGroups) != 0 {
			ExecuteTestFromTemplate(&t, resultTelevision, testRunProgress, scrollResultContainer, reportNameEntry)
		} else {
			resultTelevision.WriteAndExpand("No Test Template is imported")
		}
	})
	clearLogButton := widget.NewButton("Clear", func() {
		resultTelevision.SetText("")
	})
	importTestButton := widget.NewButton("Import", func() {
		fileDialog := dialog.NewFileOpen(
            func(r fyne.URIReadCloser, _ error) {
                // read files
				if r == nil {
					return
				}
                data, _ := ioutil.ReadAll(r)
				t, _ = CreateTestFromBytes(data)
				templateFileSelectedLabelIndicator.SetText("Template imported")
            }, w)
        fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
        fileDialog.Show()
	})
	runTestPage := container.New(	layout.NewGridLayoutWithRows(2), 
									container.New(	layout.NewHBoxLayout(), 
													container.New(	layout.NewVBoxLayout(), 
																	importTestButton, 
																	runTestButton,
																	clearLogButton,
																	layout.NewSpacer()),
													container.New(	layout.NewVBoxLayout(),
																	templateFileSelectedLabelIndicator,
																	testRunProgress,
																	reportNameEntry,
																	layout.NewSpacer(),
																),
													layout.NewSpacer()),
													scrollResultContainer,
								)
	return runTestPage
}