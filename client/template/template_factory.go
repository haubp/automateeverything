package template

import (
	"fmt"
	"os"
	"io/ioutil"
	"image/color"
	"encoding/json"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
)

var SelectedTestCase *TestCase

// CreateTestFromJSONFile Creat Test Struct from json file
func CreateTestFromJSONFile(templatePath string) (TestCategory, error) {
	templateFile, err := os.Open(templatePath)
	if err != nil {
    	fmt.Println(err)
		return TestCategory{}, err
	}
	fmt.Println("Successfully Opened ", templatePath)
	defer templateFile.Close()

    templateByteValue, _ := ioutil.ReadAll(templateFile)

	var testTemplate TestCategory
	json.Unmarshal(templateByteValue, &testTemplate)

	return testTemplate, nil
}

// CreateTestFromBytes Creat Test Struct from json file
func CreateTestFromBytes(b []byte) (TestCategory, error) {
	var testTemplate TestCategory
	json.Unmarshal(b, &testTemplate)

	return testTemplate, nil
}

// CreateJSONFileFromTemplate Creat json file from Test Struct
func CreateJSONFileFromTemplate(templatePath string, test TestCategory) {
	file, _ := json.MarshalIndent(test, "", " ")
	_ = ioutil.WriteFile(templatePath, file, 0644)
}

// CreateTestPage Creat make Test Page widget
func CreateTestPage(a fyne.App, w fyne.Window, t *TestCategory) *fyne.Container {
	manualAction := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabel("  Manual Action"),
		layout.NewSpacer(),
		widget.NewButton("+", func() {
			newW := a.NewWindow("Recording...")
			content := container.New(	layout.NewGridLayoutWithRows(3),
										layout.NewSpacer(),
										widget.NewLabel("Recording..."),
										layout.NewSpacer(),
									)
			newW.SetContent(content)
			newW.CenterOnScreen()
			newW.Resize(fyne.NewSize(300, 100))
			newW.Show()

			sl := MakeSteps()
			SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, sl...)

			newW.Close()
		}),
	)
	checkLog := container.New(
		layout.NewHBoxLayout(), 
		widget.NewLabel("  Check Log"),
		layout.NewSpacer(),
		widget.NewButton("+", func(){
			checkLogConfigWindow := a.NewWindow("Check Log Configuration")
			logPathEntry := widget.NewEntry()
			logPattern := widget.NewEntry()
			content := container.New(
				layout.NewGridLayoutWithRows(4),
				container.New(layout.NewGridLayoutWithColumns(2), widget.NewLabel("Path"), logPathEntry),
				container.New(layout.NewGridLayoutWithColumns(2), widget.NewLabel("pattern"), logPattern),
				layout.NewSpacer(),
				container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("Finish", func(){
					SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, 
						Step{
							StepName: "Check Log",
							StepAction: "CheckLog",
							StepParams: []interface{}{logPathEntry.Text, logPattern.Text},
							PreSleep: 1000,
							PostSleep: 1000,
						},
					)

					checkLogConfigWindow.Close()
				})),
			)
			checkLogConfigWindow.SetContent(content)
			checkLogConfigWindow.Resize(fyne.NewSize(400, 50))
			checkLogConfigWindow.CenterOnScreen()
			checkLogConfigWindow.Show()
		}))
	scanScreen := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabel("  Scan Screen"),
		layout.NewSpacer(),
		widget.NewButton("+", func(){
			scanScreenConfigWindow := a.NewWindow("Scan Screen Configuration")
			imgPathEntry := widget.NewEntry()
			content := container.New(
				layout.NewGridLayoutWithRows(3),
				container.New(layout.NewGridLayoutWithColumns(2), widget.NewLabel("Image Path"), imgPathEntry),
				layout.NewSpacer(),
				container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("Finish", func(){
					SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, 
						Step{
							StepName: "Check Log",
							StepAction: "CheckLog",
							StepParams: []interface{}{imgPathEntry.Text},
							PreSleep: 1000,
							PostSleep: 1000,
						},
					)
					scanScreenConfigWindow.Close()
				})),
			)
			scanScreenConfigWindow.SetContent(content)
			scanScreenConfigWindow.Resize(fyne.NewSize(400, 50))
			scanScreenConfigWindow.CenterOnScreen()
			scanScreenConfigWindow.Show()
		}),
	)

	var displayWidgets []fyne.CanvasObject

	displayWidgets = append(displayWidgets, t.GetWidget())
	for groupIndex := range t.TestCategoryGroups {
		displayWidgets = append(displayWidgets, t.TestCategoryGroups[groupIndex].GetWidget())
		for testIndex := range t.TestCategoryGroups[groupIndex].TestGroupTestCases {
			displayWidgets = append(displayWidgets, t.TestCategoryGroups[groupIndex].TestGroupTestCases[testIndex].GetWidget())
		}
	}

	var displayStepWidgets []fyne.CanvasObject
	if SelectedTestCase != nil {
		for stepIndex := range SelectedTestCase.TestCaseSteps {
			displayStepWidgets = append(displayStepWidgets, SelectedTestCase.TestCaseSteps[stepIndex].GetWidget())
		}
		fmt.Println(displayStepWidgets)
	}

	addNewStepButton := widget.NewButton("Add New Step", func() {
		newActionWindow := a.NewWindow("Recording...")
		content := container.New(layout.NewMaxLayout(), widget.NewCard("", "Action:", container.New(
			layout.NewGridLayoutWithRows(5),
			manualAction,
			checkLog,
			scanScreen,
			layout.NewSpacer(),
			container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("Add", func(){
				newActionWindow.Close()
				UpdateUI(a, w, t)
			})),
		)))
		newActionWindow.SetContent(content)
		newActionWindow.Resize(fyne.NewSize(400, 50))
		newActionWindow.CenterOnScreen()
		newActionWindow.Show()
	})

	c1 := container.NewVScroll(
        container.NewVBox(
            displayWidgets...
        ),
    )

	c2 := container.NewVScroll(
        container.NewVBox(
			displayStepWidgets...
		),
    )

	b := widget.NewButton("Export", func(){
		newW := a.NewWindow("Export Configuration")
			templatePathEntry := widget.NewEntry()
			content := container.New(	layout.NewGridLayoutWithRows(2),
										container.New(	layout.NewVBoxLayout(),
														templatePathEntry,
														layout.NewSpacer(),
										),
										container.New(	layout.NewVBoxLayout(),
														layout.NewSpacer(),
														widget.NewButton("Save", func(){
															CreateJSONFileFromTemplate(templatePathEntry.Text, *t)
															newW.Close()
														}),
										),
									)
			newW.SetContent(content)
			newW.CenterOnScreen()
			newW.Resize(fyne.NewSize(300, 100))
			newW.Show()
	})

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 1
	line.StrokeColor = color.White
	line.Position1 = fyne.NewPos(0, 0)
	line.Position2 = fyne.NewPos(0, 100)

	testPageContainer := container.New(
		layout.NewGridLayoutWithRows(2), 
		container.New(	layout.NewGridLayoutWithColumns(2), 
						c1,
						container.New(	layout.NewGridLayoutWithRows(2), 
										container.New(	layout.NewHBoxLayout(), 
														line, 
														c2),
										container.New(	layout.NewHBoxLayout(), 
														layout.NewSpacer(),
														container.New(	layout.NewVBoxLayout(), 
																		addNewStepButton,
																		layout.NewSpacer()),
										),
						),
					),
		container.New(	layout.NewVBoxLayout(), 
						layout.NewSpacer(), 
						container.New(	layout.NewHBoxLayout(),
										canvas.NewText("Category", color.RGBA{0xD8, 0xD8, 0xD8, 1}),
										canvas.NewText("Group", color.RGBA{0x56, 0x2B, 0x08, 1}),
										canvas.NewText("Test Case", color.RGBA{0x64, 0x7E, 0x68, 1}),
										layout.NewSpacer(),
										b,
									),
					),
	)

	return testPageContainer
}

// CreateRunTestPage Creat Run Test Page widget
func CreateRunTestPage(a fyne.App, w fyne.Window, t *TestCategory) *fyne.Container {
	importTestButton := widget.NewButton("Import", func() {
		fmt.Println("Import clicked")
	})
	testProgress := widget.NewProgressBar()
	runTestButton := widget.NewButton("Run", func() {
		fmt.Println("tapped")
	})
	runTestPage := container.New(layout.NewGridLayoutWithRows(3), importTestButton, testProgress, runTestButton)

	return runTestPage
}

// UpdateUI update test page UI
func UpdateUI(a fyne.App, w fyne.Window, t *TestCategory) {
	tabs := container.NewAppTabs(
		container.NewTabItem("Create Test", CreateTestPage(a, w, t)),
		container.NewTabItem("Run Test", CreateRunTestPage(a, w, t)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.CenterOnScreen()
	w.SetContent(tabs)
}