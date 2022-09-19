package main

import (
	"automateeverything.com/v2/template"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"io/ioutil"
)

func createMenu(a fyne.App, w fyne.Window, t *template.TestCategory) *fyne.MainMenu {
	importItem := fyne.NewMenuItem("Import", func() {
		// Using dialogs to open files
        // first argument func(fyne.URIReadCloser, error)
        // 2nd is parent window in our case "w"
        // r for reader
        // _ is ignore error
        fileDialog := dialog.NewFileOpen(
            func(r fyne.URIReadCloser, _ error) {
				template.SelectedTestCase = nil
                // read files
                data, _ := ioutil.ReadAll(r)
				tt, _ := template.CreateTestFromBytes(data)
				*t = tt
				t.InitContext(a, w)
				for groupIndex := range t.TestCategoryGroups {
					t.TestCategoryGroups[groupIndex].InitContext(a, w, t)
					for testCaseIndex := range t.TestCategoryGroups[groupIndex].TestGroupTestCases {
						t.TestCategoryGroups[groupIndex].TestGroupTestCases[testCaseIndex].InitContext(a, w, t)
						for testStepIndex := range t.TestCategoryGroups[groupIndex].TestGroupTestCases[testCaseIndex].TestCaseSteps {
							t.TestCategoryGroups[groupIndex].TestGroupTestCases[testCaseIndex].TestCaseSteps[testStepIndex].InitContext(a, w)
						}
					}
				}
				tabs := container.NewAppTabs(
					container.NewTabItem("Create Test", template.CreateTestPage(a, w, t)),
					container.NewTabItem("Run Test", template.CreateRunTestPage(a, w)),
				)
				tabs.SetTabLocation(container.TabLocationLeading)
			
				w.CenterOnScreen()
				w.SetContent(tabs)
            }, w)
        fileDialog.SetFilter(
            storage.NewExtensionFileFilter([]string{".json"}))
        fileDialog.Show()
	})

	fileMenu := fyne.NewMenu("File", importItem)
	menu := fyne.NewMainMenu(fileMenu)

	return menu
}

// TODO: check decrypt log, enter text from keyboard action, enhance UX

func main() {
	// runner.ExecuteTest("myTemplate.json")

	// steps := template.MakeSteps()
	// test := template.Test{TestName:"any", TestSteps: steps}
	// group := template.TestGroup{TestGroupName:"any", TestGroupTests: []template.Test{test}}
	// category := template.TestCategory{TestCategoryName:"any", TestCategoryGroups: []template.TestGroup{group}}
	// template.CreateJsonFileFromTemplate("myTemplate2.json", category)

	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	w := myApp.NewWindow("Automation Testing")
	w.Resize(fyne.NewSize(1100, 800))

	// Test data
	t, _ := template.CreateTestFromJSONFile("template.json")
	t.InitContext(myApp, w)
	for groupIndex := range t.TestCategoryGroups {
		t.TestCategoryGroups[groupIndex].InitContext(myApp, w, &t)
		for testCaseIndex := range t.TestCategoryGroups[groupIndex].TestGroupTestCases {
			t.TestCategoryGroups[groupIndex].TestGroupTestCases[testCaseIndex].InitContext(myApp, w, &t)
			for testStepIndex := range t.TestCategoryGroups[groupIndex].TestGroupTestCases[testCaseIndex].TestCaseSteps {
				t.TestCategoryGroups[groupIndex].TestGroupTestCases[testCaseIndex].TestCaseSteps[testStepIndex].InitContext(myApp, w)
			}
		}
	}

	// Menu
	w.SetMainMenu(createMenu(myApp, w, &t))

	// Tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Create Test", template.CreateTestPage(myApp, w, &t)),
		container.NewTabItem("Run Test", template.CreateRunTestPage(myApp, w)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.CenterOnScreen()
	w.SetContent(tabs)
	w.Show()
	myApp.Run()
}
