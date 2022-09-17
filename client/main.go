package main

import (
	// "automateeverything.com/v2/runner"
	"automateeverything.com/v2/template"
	"automateeverything.com/v2/cwidget"

	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
)

func createMenu() *fyne.MainMenu {
	importItem := fyne.NewMenuItem("Import", func() {
		fmt.Println("Import pressed")
	})
	fileMenu := fyne.NewMenu("File", importItem)
	menu := fyne.NewMainMenu(fileMenu)

	return menu
}

func createTestPage(a fyne.App, t *template.TestCategory) *fyne.Container {
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

	// var testGroupWidgets []cWidget.TestWidget
	// var testCaseWidgets []cWidget.TestWidget

	var displayWidget []fyne.CanvasObject
	var referenceWidget []cwidget.TestWidget

	displayWidget = append(displayWidget, widget.NewLabel(t.TestCategoryName))
	referenceWidget = append(referenceWidget, cwidget.NewCategoryWidget(t))
	
	for _, group := range t.TestCategoryGroups {
		displayWidget = append(displayWidget, widget.NewLabel(group.TestGroupName))
		
		for _, test := range group.TestGroupTestCases {
			displayWidget = append(displayWidget, widget.NewLabel(test.TestName))

		}
	}

	t4 := widget.NewButton("Add New Step", func() {
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
        container.NewVBox(
            t4,
        ),
    )

	b := widget.NewButton("Export", func(){})

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 1
	line.StrokeColor = color.White
	line.Position1 = fyne.NewPos(0, 0)
	line.Position2 = fyne.NewPos(0, 100)

	createTestPage := container.New(
		layout.NewGridLayoutWithRows(2), 
		container.New(	layout.NewGridLayoutWithColumns(2), 
						c1,
						container.New(layout.NewHBoxLayout(), line, c2),
						),
		container.New(layout.NewVBoxLayout(), layout.NewSpacer(), container.New(layout.NewHBoxLayout(), layout.NewSpacer(), b)))

	return createTestPage
}

func main() {
	// runner.ExecuteTest("myTemplate.json")

	// steps := template.MakeSteps()
	// test := template.Test{TestName:"any", TestSteps: steps}
	// group := template.TestGroup{TestGroupName:"any", TestGroupTests: []template.Test{test}}
	// category := template.TestCategory{TestCategoryName:"any", TestCategoryGroups: []template.TestGroup{group}}
	// template.CreateJsonFileFromTemplate("myTemplate2.json", category)

	t, _ := template.CreateTestFromJsonFile("myTemplate.json")
	fmt.Println(t)

	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())
	w := myApp.NewWindow("Automation Testing")
	w.Resize(fyne.NewSize(800, 800))

	// Menu
	w.SetMainMenu(createMenu())

	// Create Test Page
	createTestPage := createTestPage(myApp, &t)

	// Run Test Page
	importTestButton := widget.NewButton("Import", func() {
		fmt.Println("Import clicked")
	})
	testProgress := widget.NewProgressBar()
	runTestButton := widget.NewButton("Run", func() {
		fmt.Println("tapped")
	})
	runTestPage := container.New(layout.NewGridLayoutWithRows(3), importTestButton, testProgress, runTestButton)

	// Tabs
	tabs := container.NewAppTabs(
		container.NewTabItem("Create Test", createTestPage),
		container.NewTabItem("Run Test", runTestPage),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	w.CenterOnScreen()
	w.SetContent(tabs)
	w.ShowAndRun()
}
