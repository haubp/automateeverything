package template

import (
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"image/color"
	"encoding/json"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2"

	"automateeverything.com/v2/utils"
)

// SelectedTestCase Global Selected Test Case => TODO: Need to remove global variable
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
	file, e := json.MarshalIndent(test, "", " ")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(test)
	e = ioutil.WriteFile(templatePath, file, 0644)
	if e != nil {
		fmt.Println(e)
	}
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
			for i := range sl {
				sl[i].InitContext(a, w)
			}

			fmt.Printf("selected now point to %p\n", SelectedTestCase)

			SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, sl...)

			fmt.Printf("t test case %p", &t.TestCategoryGroups[0].TestGroupTestCases[0])

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
					newStep := Step{
						StepName: "Check Log",
						StepAction: "CheckLog",
						StepParams: []interface{}{logPathEntry.Text, logPattern.Text},
						PreSleep: 1000,
						PostSleep: 1000,
					}
					newStep.InitContext(a, w)

					SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, newStep)

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
					newStep := Step{
						StepName: "Check Log",
						StepAction: "CheckLog",
						StepParams: []interface{}{imgPathEntry.Text},
						PreSleep: 1000,
						PostSleep: 1000,
					}
					newStep.InitContext(a, w)
					SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, newStep)
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

	selectedTestCaseIndicator := widget.NewLabel("")
	if SelectedTestCase != nil {
		selectedTestCaseIndicator.SetText(SelectedTestCase.TestCaseName + " selected: ")
	}
	
	testPageContainer := container.New(
		layout.NewGridLayoutWithRows(2), 
		container.New(	layout.NewGridLayoutWithColumns(2), 
						c1,
						container.New(	layout.NewGridLayoutWithRows(2), 
										container.New(	layout.NewHBoxLayout(), 
														line, 
														selectedTestCaseIndicator,
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
func CreateRunTestPage(a fyne.App, w fyne.Window) *fyne.Container {
	var t TestCategory
	runTestButton := widget.NewButton("Run", func() {
		ExecuteTestFromTemplate(&t)
	})
	templateFileSelectedLabelIndicator := widget.NewLabel("Template not selected")
	importTestButton := widget.NewButton("Import", func() {
		fileDialog := dialog.NewFileOpen(
            func(r fyne.URIReadCloser, _ error) {
                // read files
                data, _ := ioutil.ReadAll(r)
				t, _ = CreateTestFromBytes(data)
				templateFileSelectedLabelIndicator.SetText("Template selected")
            }, w)
        fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))
        fileDialog.Show()
	})
	runTestPage := container.New(	layout.NewGridLayoutWithRows(3), 
									container.New(	layout.NewVBoxLayout(), 
													layout.NewSpacer(), 
													container.New(	layout.NewHBoxLayout(), 
																	importTestButton, 
																	layout.NewSpacer()),
													layout.NewSpacer()),
									container.New(	layout.NewVBoxLayout(), 
													layout.NewSpacer(), 
													container.New(	layout.NewHBoxLayout(), 
																	templateFileSelectedLabelIndicator, 
																	layout.NewSpacer()),
													layout.NewSpacer()),
									container.New(	layout.NewVBoxLayout(), 
													layout.NewSpacer(), 
													container.New(	layout.NewHBoxLayout(), 
																	layout.NewSpacer(),
																	runTestButton,
																	layout.NewSpacer()),
													layout.NewSpacer()),
								)
	return runTestPage
}

// UpdateUI update test page UI
func UpdateUI(a fyne.App, w fyne.Window, t *TestCategory) {
	tabs := container.NewAppTabs(
		container.NewTabItem("Create Test", CreateTestPage(a, w, t)),
		container.NewTabItem("Run Test", CreateRunTestPage(a, w)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.CenterOnScreen()
	w.SetContent(tabs)
}

// ExecuteTestFromTemplatePath execute test from file path
func ExecuteTestFromTemplatePath(templatePath string) {
	actionsMap := utils.ActionsMap()

	testTemplate, err := CreateTestFromJSONFile(templatePath)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	fmt.Println("Run test for category", testTemplate.TestCategoryName)
	for _, group := range testTemplate.TestCategoryGroups {
		fmt.Println("	Run test for group", group.TestGroupName)
		for _, test := range group.TestGroupTestCases {
			fmt.Println("		Run test for test", test.TestCaseName)
			capturedTime := time.Now()
			for _, step := range test.TestCaseSteps {
				fmt.Println("			Run step", step.StepName)
				if step.StepAction == "CaptureTime" {
					capturedTime = time.Now()
				} else {
					if step.StepAction == "CheckLog" {
						step.StepParams = append(step.StepParams, capturedTime)
					}

					time.Sleep(time.Duration(step.PreSleep) * time.Millisecond)
					if actionsMap[step.StepAction].(func([]interface{}) bool)(step.StepParams) {
						fmt.Println("				", "Success")
					} else {
						fmt.Println("				", "Failed")
						break
					}
					time.Sleep(time.Duration(step.PostSleep) * time.Millisecond)
				}
			}
		}
	}
}

// ExecuteTestFromTemplate execute test from file bytes
func ExecuteTestFromTemplate(t *TestCategory) {
	actionsMap := utils.ActionsMap()

	fmt.Println("Run test for category", t.TestCategoryName)
	for _, group := range t.TestCategoryGroups {
		fmt.Println("	Run test for group", group.TestGroupName)
		for _, test := range group.TestGroupTestCases {
			fmt.Println("		Run test for test", test.TestCaseName)
			capturedTime := time.Now()
			for _, step := range test.TestCaseSteps {
				fmt.Println("			Run step", step.StepName)
				if step.StepAction == "CaptureTime" {
					capturedTime = time.Now()
				} else {
					if step.StepAction == "CheckLog" {
						step.StepParams = append(step.StepParams, capturedTime)
					}

					time.Sleep(time.Duration(step.PreSleep) * time.Millisecond)
					if actionsMap[step.StepAction].(func([]interface{}) bool)(step.StepParams) {
						fmt.Println("				", "Success")
					} else {
						fmt.Println("				", "Failed")
						break
					}
					time.Sleep(time.Duration(step.PostSleep) * time.Millisecond)
				}
			}
		}
	}
}