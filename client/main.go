package main

import (
	// "automateeverything.com/v1/runner"
	// "automateeverything.com/v1/template"

	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// runner.ExecuteTest("myTemplate.json")

	// steps := template.MakeSteps()
	// test := template.Test{TestName:"any", TestSteps: steps}
	// group := template.TestGroup{TestGroupName:"any", TestGroupTests: []template.Test{test}}
	// category := template.TestCategory{TestCategoryName:"any", TestCategoryGroups: []template.TestGroup{group}}
	// template.CreateJsonFileFromTemplate("myTemplate2.json", category)

	myApp := app.New()
	w := myApp.NewWindow("Automation Testing")
	w.Resize(fyne.NewSize(800, 800))

	// Menu
	importItem := fyne.NewMenuItem("Import", func() {
		fmt.Println("Import pressed")
	})
	fileMenu := fyne.NewMenu("File", importItem)
	menu := fyne.NewMainMenu(fileMenu)
	w.SetMainMenu(menu)

	// Create Test Page

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
			checkLogConfigWindow := myApp.NewWindow("Check Log Configuration")
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
			scanScreenConfigWindow := myApp.NewWindow("Scan Screen Configuration")
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

	t1 := widget.NewLabel("Test Category")
	t2 := widget.NewLabel("Test Groups")
	t3 := widget.NewLabel("Test Cases")

	t4 := widget.NewButton("Add New Step", func() {
		newActionWindow := myApp.NewWindow("New Action")
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
            t1,
			t2,
			t3,
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
