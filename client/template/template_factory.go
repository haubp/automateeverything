package template

import (
	"encoding/json"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"

	"automateeverything.com/v2/custom"
	"automateeverything.com/v2/utils"
	// "automateeverything.com/v2/custom"
)

// SelectedTestCase Global Selected Test Case => TODO: Need to remove global variable
var SelectedTestCase *TestCase

// CreateTestFromJSONFile Creat Test Struct from json file
func CreateTestFromJSONFile(templatePath string) (TestCategory, error) {
	templateFile, err := os.Open(templatePath)
	if err != nil {
    	log.Println(err)
		return TestCategory{}, err
	}
	log.Println("Successfully Opened ", templatePath)
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
		log.Println(e)
	}
	e = ioutil.WriteFile(templatePath, file, 0644)
	if e != nil {
		log.Println(e)
	}
}

// CreateTestPage Creat make Test Page widget
func CreateTestPage(a fyne.App, w fyne.Window, t *TestCategory) *fyne.Container {
	manualAction := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabel("  Manual Action"),
		layout.NewSpacer(),
		widget.NewButton("+", func() {
			newW := a.NewWindow("Manual Action")
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

			SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, sl...)

			UpdateUI(a, w, t)

			newW.Close()
		}),
	)
	keyboardTyping := container.New(
		layout.NewHBoxLayout(), 
		widget.NewLabel("  Keyboard Action"),
		layout.NewSpacer(),
		widget.NewButton("+", func(){
			keyboardTypingConfigWindow := a.NewWindow("keyboard Typing Configuration")
			stringEntry := widget.NewEntry()
			content := container.New(
				layout.NewGridLayoutWithRows(4),
				container.New(layout.NewGridLayoutWithColumns(2), widget.NewLabel("text"), stringEntry),
				layout.NewSpacer(),
				container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("Finish", func(){
					newStep := Step{
						StepName: "Keyboard Typing",
						StepAction: "KeyboardTyping",
						StepParams: []interface{}{stringEntry.Text},
						PreSleep: 1000,
						PostSleep: 1000,
					}
					newStep.InitContext(a, w)

					SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, newStep)

					UpdateUI(a, w, t)

					keyboardTypingConfigWindow.Close()
				})),
			)
			keyboardTypingConfigWindow.SetContent(content)
			keyboardTypingConfigWindow.Resize(fyne.NewSize(400, 50))
			keyboardTypingConfigWindow.CenterOnScreen()
			keyboardTypingConfigWindow.Show()
		}))
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

					UpdateUI(a, w, t)

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

					UpdateUI(a, w, t)

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
		newActionWindow := a.NewWindow("Add New Step")
		content := container.New(layout.NewMaxLayout(), widget.NewCard("", "Action:", container.New(
			layout.NewGridLayoutWithRows(6),
			manualAction,
			checkLog,
			scanScreen,
			keyboardTyping,
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
			newW.Canvas().Focus(templatePathEntry)
	})

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 1
	line.StrokeColor = color.Gray{10}
	line.Position1 = fyne.NewPos(0, 0)

	selectedTestCaseIndicator := canvas.NewText("", color.RGBA{0xD8, 0xD8, 0xD8, 1})
	if SelectedTestCase != nil {
		selectedTestCaseIndicator.Text = "\t" + SelectedTestCase.TestCaseName + " selected"
	} else {
		selectedTestCaseIndicator.Text = "\t" + "No Test Case Selected"
	}
	
	testPageContainer := container.New(
		layout.NewGridLayoutWithRows(2), 
		container.New(	layout.NewGridLayoutWithColumns(2), 
						c1,
						container.New(	layout.NewGridLayoutWithRows(2),
										container.New(	layout.NewHBoxLayout(), 
														line, 
														c2),
										container.New(	layout.NewHBoxLayout(),
														container.New(	layout.NewVBoxLayout(),
																		selectedTestCaseIndicator,
																		layout.NewSpacer()),
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

// RunTestPage Creat Run Test Page widget
func RunTestPage(a fyne.App, w fyne.Window) *fyne.Container {
	var t TestCategory
	resultTelevision := custom.NewFixSizeLabel()
	resultTelevision.SetFixSize(fyne.NewSize(500, 100))
	testRunProgress := widget.NewProgressBar()
	testRunProgress.SetValue(0.0)
	templateFileSelectedLabelIndicator := widget.NewLabel("No test template found")
	scrollResultContainer := container.NewVScroll(
		resultTelevision,
	)
	runTestButton := widget.NewButton("Run", func() {
		if len(t.TestCategoryGroups) != 0 {
			ExecuteTestFromTemplate(&t, resultTelevision, testRunProgress, scrollResultContainer)
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
																	layout.NewSpacer(),
																),
													layout.NewSpacer()),
													scrollResultContainer,
								)
	return runTestPage
}

// UpdateUI update test page UI
func UpdateUI(a fyne.App, w fyne.Window, t *TestCategory) {
	tabs := container.NewAppTabs(
		container.NewTabItem("Create Test", CreateTestPage(a, w, t)),
		container.NewTabItem("Run Test", RunTestPage(a, w)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.CenterOnScreen()
	w.SetContent(tabs)
}

// ExecuteTestFromTemplate execute test from file bytes
func ExecuteTestFromTemplate(t *TestCategory, resultTelevision *custom.FixSizeLabel, progressBar *widget.ProgressBar, resultContainer *container.Scroll) {
	actionsMap := utils.ActionsMap()
	resultTelevision.WriteAndExpand("Run test for category " + t.TestCategoryName)
	totalOfTestCases := 0.0
	for _, group := range t.TestCategoryGroups {
		totalOfTestCases += float64(len(group.TestGroupTestCases))
	}
	percentOfEachTestCase := 100 / totalOfTestCases
	for _, group := range t.TestCategoryGroups {
		resultTelevision.WriteAndExpand("	Run test for group " + group.TestGroupName)
		for _, test := range group.TestGroupTestCases {
			resultTelevision.WriteAndExpand("		Run test for test " + test.TestCaseName)
			var capturedTime time.Time
			for _, step := range test.TestCaseSteps {
				resultTelevision.WriteAndExpand("			Run step " + step.StepName)
				if step.StepAction == "CaptureTime" {
					capturedTime = time.Now()
					resultTelevision.WriteAndExpand("			Capture time")
				} else {
					if step.StepAction == "CheckLog" {
						step.StepParams = append(step.StepParams, capturedTime)
					}
					time.Sleep(time.Duration(step.PreSleep) * time.Millisecond)
					if actionsMap[step.StepAction].(func([]interface{}) bool)(step.StepParams) {
						resultTelevision.WriteAndExpand("				" + "Success")
					} else {
						resultTelevision.WriteAndExpand("				" + "Failed")
						break
					}
					time.Sleep(time.Duration(step.PostSleep) * time.Millisecond)
				}
				resultContainer.ScrollToBottom()
			}
			progressBar.SetValue(progressBar.Value + percentOfEachTestCase)
		}
	}
	resultTelevision.WriteAndExpand("Finish")
}
