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
func CreateTestPage(a fyne.App, t *TestCategory) *fyne.Container {
	captureTime := container.New(
		layout.NewHBoxLayout(),
		widget.NewLabel("Capture Time"),
		layout.NewSpacer(),
		widget.NewButton("+", func(){}),
	)
	checkLog := container.New(
		layout.NewHBoxLayout(), 
		widget.NewLabel("Check Log"),
		layout.NewSpacer(),
		widget.NewButton("+", func(){
			checkLogConfigWindow := a.NewWindow("Check Log Configuration")
			content := container.New(
				layout.NewGridLayoutWithRows(4),
				container.New(layout.NewGridLayoutWithColumns(2), widget.NewLabel("Path"), widget.NewEntry()),
				container.New(layout.NewGridLayoutWithColumns(2), widget.NewLabel("pattern"), widget.NewEntry()),
				layout.NewSpacer(),
				container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("Finish", func(){
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
		widget.NewLabel("Scan Screen"),
		layout.NewSpacer(),
		widget.NewButton("+", func(){
			scanScreenConfigWindow := a.NewWindow("Scan Screen Configuration")
			content := container.New(
				layout.NewGridLayoutWithRows(3),
				container.New(layout.NewGridLayoutWithColumns(2), widget.NewLabel("Image Path"), widget.NewEntry()),
				layout.NewSpacer(),
				container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("Finish", func(){
					scanScreenConfigWindow.Close()
				})),
			)
			scanScreenConfigWindow.SetContent(content)
			scanScreenConfigWindow.Resize(fyne.NewSize(400, 50))
			scanScreenConfigWindow.CenterOnScreen()
			scanScreenConfigWindow.Show()
		}),
	)

	var displayWidget []fyne.CanvasObject

	displayWidget = append(displayWidget, t.GetWidget())
	
	for groupIndex := range t.TestCategoryGroups {
		displayWidget = append(displayWidget, t.TestCategoryGroups[groupIndex].GetWidget())
		for testIndex := range t.TestCategoryGroups[groupIndex].TestGroupTestCases {
			displayWidget = append(displayWidget, t.TestCategoryGroups[groupIndex].TestGroupTestCases[testIndex].GetWidget())
		}
	}

	addNewStepButton := widget.NewButton("Add New Step", func() {
		newActionWindow := a.NewWindow("New Action")
		content := container.New(layout.NewMaxLayout(), widget.NewCard("Step Action", "", container.New(
			layout.NewGridLayoutWithRows(5),
			captureTime,
			checkLog,
			scanScreen,
			layout.NewSpacer(),
			container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("Add", func(){
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
            displayWidget...
        ),
    )

	c2 := container.NewVScroll(
        container.NewVBox(),
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
						container.New(	layout.NewVBoxLayout(), 
										container.New(layout.NewHBoxLayout(), 
										line, 
										c2, 
										layout.NewSpacer()), 
										container.New(layout.NewHBoxLayout(), layout.NewSpacer(), addNewStepButton),
						),
					),
		container.New(	layout.NewVBoxLayout(), 
						layout.NewSpacer(), 
						container.New(	layout.NewHBoxLayout(), 
										canvas.NewText("Category", color.RGBA{100, 0, 0, 1}),
										canvas.NewText("Group", color.RGBA{0, 0, 100, 1}),
										canvas.NewText("Test Case", color.RGBA{0, 100, 0, 1}),
										layout.NewSpacer(),
										b,
									),
					),
	)

	return testPageContainer
}

// CreateRunTestPage Creat Run Test Page widget
func CreateRunTestPage(a fyne.App, t *TestCategory) *fyne.Container {
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