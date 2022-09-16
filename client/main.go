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

	a := app.New()
	w := a.NewWindow("Automation Testing")
	w.Resize(fyne.NewSize(800, 800))

	// Menu
	importItem := fyne.NewMenuItem("Import", func() {
		fmt.Println("Import pressed")
	})
	fileMenu := fyne.NewMenu("File", importItem)
	menu := fyne.NewMainMenu(fileMenu)
	w.SetMainMenu(menu)

	// Create Test Page
	exportButton := widget.NewButton("Export", func() {
		fmt.Println("Export clicked")
	})
	exportButton.Resize(fyne.NewSize(50, 30))

	tests := widget.NewCard("tests", "tests", canvas.NewText("tests", color.White))
	steps := widget.NewCard("steps", "steps", canvas.NewText("steps", color.White))
	makesteps := widget.NewCard("makesteps", "makesteps", canvas.NewText("makesteps", color.White))

	createTestPage := container.New(
		layout.NewVBoxLayout(),
		container.New(
			layout.NewGridLayoutWithColumns(2), 
			container.New(layout.NewGridLayoutWithRows(1), tests),
			container.New(layout.NewGridLayoutWithRows(2), steps, makesteps),
		),
		layout.NewSpacer(),
		container.New(layout.NewHBoxLayout(), layout.NewSpacer(), exportButton, layout.NewSpacer()),
	)
	
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

	w.SetContent(tabs)
	w.ShowAndRun()
}
