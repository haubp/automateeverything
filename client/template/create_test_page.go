package template

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateTestPage Create make Test Page widget
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

			if SelectedTestCase != nil {
				SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, sl...)
			}
			UpdateUI(a, w, t)

			newW.Close()
		}),
	)
	runCommandAction := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabel("  Run Command Action"),
		layout.NewSpacer(),
		widget.NewButton("+", func() {
			runCommandConfigWindow := a.NewWindow("Command Configuration")
			commandEntry := widget.NewEntry()
			content := container.New(
				layout.NewGridLayoutWithRows(2),
				container.New(layout.NewGridLayoutWithColumns(2), widget.NewLabel("Command"), commandEntry),
				layout.NewSpacer(),
				container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("Finish", func(){
					newStep := Step{
						StepName: "Run Command",
						StepAction: "RunCommand",
						StepParams: []interface{}{commandEntry.Text},
						PreSleep: 1000,
						PostSleep: 1000,
					}
					newStep.InitContext(a, w)

					if SelectedTestCase != nil {
						SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, &newStep)
					}

					UpdateUI(a, w, t)

					runCommandConfigWindow.Close()
				})),
			)
			runCommandConfigWindow.SetContent(content)
			runCommandConfigWindow.Resize(fyne.NewSize(400, 50))
			runCommandConfigWindow.CenterOnScreen()
			runCommandConfigWindow.Canvas().Focus(commandEntry)
			runCommandConfigWindow.Show()
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

					if SelectedTestCase != nil {
						SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, &newStep)
					}

					UpdateUI(a, w, t)

					keyboardTypingConfigWindow.Close()
				})),
			)
			keyboardTypingConfigWindow.SetContent(content)
			keyboardTypingConfigWindow.Resize(fyne.NewSize(400, 50))
			keyboardTypingConfigWindow.CenterOnScreen()
			keyboardTypingConfigWindow.Canvas().Focus(stringEntry)
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

					if SelectedTestCase != nil {
						SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, &newStep)
					}

					UpdateUI(a, w, t)

					checkLogConfigWindow.Close()
				})),
			)
			checkLogConfigWindow.SetContent(content)
			checkLogConfigWindow.Resize(fyne.NewSize(400, 50))
			checkLogConfigWindow.CenterOnScreen()
			checkLogConfigWindow.Canvas().Focus(logPathEntry)
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

					if SelectedTestCase != nil {
						SelectedTestCase.TestCaseSteps = append(SelectedTestCase.TestCaseSteps, &newStep)
					}

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

	upperBorder := canvas.NewRectangle(color.White)
	upperBorder.SetMinSize(fyne.NewSize(500, 1))
	displayStepWidgets := []fyne.CanvasObject{ upperBorder }
	if SelectedTestCase != nil {
		for stepIndex := range SelectedTestCase.TestCaseSteps {
			displayStepWidgets = append(displayStepWidgets, SelectedTestCase.TestCaseSteps[stepIndex].GetWidget())
		}
	}
	bottomBorder := canvas.NewRectangle(color.White)
	bottomBorder.SetMinSize(fyne.NewSize(500, 1))
	displayStepWidgets = append(displayStepWidgets, bottomBorder)

	addNewStepButton := widget.NewButton("Add New Step", func() {
		newActionWindow := a.NewWindow("Add New Step")
		content := container.New(layout.NewMaxLayout(), widget.NewCard("", "Action:", container.New(
			layout.NewGridLayoutWithRows(7),
			runCommandAction,
			manualAction,
			checkLog,
			scanScreen,
			keyboardTyping,
			layout.NewSpacer(),
			container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("Done", func(){
				UpdateUI(a, w, t)
				newActionWindow.Close()
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

	rightBorder := canvas.NewRectangle(color.White)
	rightBorder.SetMinSize(fyne.NewSize(1, 300))
	
	testPageContainer := container.New(
		layout.NewGridLayoutWithRows(2), 
		container.New(	layout.NewGridLayoutWithColumns(2), 
						c1,
						container.New(	layout.NewVBoxLayout(),
										container.New(	layout.NewHBoxLayout(), 
														rightBorder,
														c2),
										layout.NewSpacer(),
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