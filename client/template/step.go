package template

import (
	"image/color"
	"log"
	"strconv"

	"automateeverything.com/v2/custom"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

// Step type
type Step struct {
	StepName   string          `json:"step_name"`
	StepAction string          `json:"step_action"`
	StepParams []interface{}   `json:"step_params"`
	PreSleep   int             `json:"step_pre_sleep"`
	PostSleep  int             `json:"step_post_sleep"`
	Widget     *fyne.Container `json:"-"`
	W          fyne.Window     `json:"-"`
	A          fyne.App        `json:"-"`
}

// MakeSteps auto
func MakeSteps() []*Step {
	steps := []*Step{}

	log.Println("Start Hook")

	hook.Register(hook.KeyDown, []string{"d"}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		log.Println("double click", x, y)
		steps = append(steps, &Step{
			StepName:   "Move Mouse",
			StepAction: "MoveMouse",
			StepParams: []interface{}{x, y},
			PreSleep:   1000,
			PostSleep:  1000})
		steps = append(steps, &Step{
			StepName:   "Double Click",
			StepAction: "LeftClick",
			StepParams: []interface{}{true},
			PreSleep:   1000,
			PostSleep:  1000})
	})

	hook.Register(hook.KeyDown, []string{"c"}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		log.Println("click", x, y)
		steps = append(steps, &Step{
			StepName:   "Move Mouse",
			StepAction: "MoveMouse",
			StepParams: []interface{}{x, y},
			PreSleep:   1000,
			PostSleep:  1000})
		steps = append(steps, &Step{
			StepName:   "Left Click",
			StepAction: "LeftClick",
			StepParams: []interface{}{false},
			PreSleep:   1000,
			PostSleep:  1000})
	})

	hook.Register(hook.KeyDown, []string{"r"}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		log.Println("right click", x, y)
		steps = append(steps, &Step{
			StepName:   "Move Mouse",
			StepAction: "MoveMouse",
			StepParams: []interface{}{x, y},
			PreSleep:   1000,
			PostSleep:  1000})
		steps = append(steps, &Step{
			StepName:   "Right Click",
			StepAction: "RightClick",
			StepParams: []interface{}{false},
			PreSleep:   1000,
			PostSleep:  1000})
	})

	hook.Register(hook.KeyDown, []string{"s"}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		log.Println("scroll", x, y)
		steps = append(steps, &Step{
			StepName:   "Move Mouse",
			StepAction: "MoveMouse",
			StepParams: []interface{}{x, y},
			PreSleep:   1000,
			PostSleep:  1000})
		steps = append(steps, &Step{
			StepName:   "Scroll Mouse",
			StepAction: "ScrollMouse",
			StepParams: []interface{}{100},
			PreSleep:   1000,
			PostSleep:  1000})
	})

	hook.Register(hook.KeyDown, []string{"t"}, func(e hook.Event) {
		steps = append(steps, &Step{
			StepName:   "Capture Time",
			StepAction: "CaptureTime",
			StepParams: []interface{}{},
			PreSleep:   100,
			PostSleep:  100})
	})

	hook.Register(hook.KeyDown, []string{"q"}, func(e hook.Event) {
		log.Println("exit")
		hook.End()
	})

	log.Println("Listening")

	s := hook.Start()
	<-hook.Process(s)

	return steps
}

func GetNewPositionOnScreen() *Step {
	var step *Step

	log.Println("Start Hook")

	hook.Register(hook.KeyDown, []string{"m"}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		log.Println("double click", x, y)
		step = &Step{
			StepName:   "Move Mouse",
			StepAction: "MoveMouse",
			StepParams: []interface{}{x, y},
			PreSleep:   1000,
			PostSleep:  1000}
	})

	hook.Register(hook.KeyDown, []string{"q"}, func(e hook.Event) {
		log.Println("exit")
		hook.End()
	})

	log.Println("Listening")

	s := hook.Start()
	<-hook.Process(s)

	return step
}

// find index of step in steps
func (s *Step) FindIndex(steps []*Step) int {
	for i, step := range steps {
		if step == s {
			return i
		}
	}
	return -1
}

// InitContext init step context
func (c *Step) InitContext(a fyne.App, w fyne.Window) {
	c.W = w
	c.A = a

	c.Widget = container.New(layout.NewHBoxLayout(),
		canvas.NewText(c.StepName, color.RGBA{0xD8, 0xD8, 0xD8, 1}),
		layout.NewSpacer(),

		widget.NewButton("-", func() {
			SelectedTestCase.DeleteSelectedStepByPointer(c)
			UpdateUI(a, w, SelectedTestCase.Category)

		}),

		widget.NewButton("+", func() {
			index := c.FindIndex(SelectedTestCase.TestCaseSteps)

			manualAction := createManualActionWidget(a, w, SelectedTestCase.Category, index)
			runCommandAction := createRunCommandActionWidget(a, w, SelectedTestCase.Category, index)
			keyboardTyping := createKeyboardTypingWidget(a, w, SelectedTestCase.Category, index)
			checkLog := createCheckLogWidget(a, w, SelectedTestCase.Category, index)
			scanScreen := createScanScreenWidget(a, w, SelectedTestCase.Category, index)
			maConfig := createMAConfigWidget(a, w, SelectedTestCase.Category, index)

			newActionWindow := a.NewWindow("Add New Step Before Current Step")
			content := container.New(layout.NewMaxLayout(), widget.NewCard("", "Action:", container.New(
				layout.NewGridLayoutWithRows(8),
				runCommandAction,
				manualAction,
				checkLog,
				scanScreen,
				keyboardTyping,
				maConfig,

				layout.NewSpacer(),
				container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("Done", func() {
					UpdateUI(a, w, SelectedTestCase.Category)
					newActionWindow.Close()
				})),
			)))
			newActionWindow.SetContent(content)
			newActionWindow.Resize(fyne.NewSize(400, 50))
			newActionWindow.CenterOnScreen()
			newActionWindow.Show()
		}),

		widget.NewButton("...", func() {
			newW := a.NewWindow("Step configuration")

			preDelayEntry := custom.NewFixSizeEntry(fyne.NewSize(100, 10))
			postDelayEntry := custom.NewFixSizeEntry(fyne.NewSize(100, 10))
			addPreStep := custom.NewFixSizeEntry(fyne.NewSize(100, 10))
			preDelayEntry.SetText(strconv.FormatInt(int64(c.PreSleep), 10))
			postDelayEntry.SetText(strconv.FormatInt(int64(c.PostSleep), 10))
			addPreStep.SetText(strconv.FormatInt(int64(c.PostSleep), 10))

			areaPreDelay := container.New(layout.NewHBoxLayout(),
				widget.NewLabel("Pre Delay"),
				preDelayEntry,
				layout.NewSpacer(),
				widget.NewButton("Update", func() {
					c.PreSleep, _ = strconv.Atoi(preDelayEntry.Text)
				}))

			areaPostDelay := container.New(layout.NewHBoxLayout(),
				widget.NewLabel("Post Delay"),
				postDelayEntry,
				layout.NewSpacer(),
				widget.NewButton("Update", func() {
					c.PostSleep, _ = strconv.Atoi(postDelayEntry.Text)
				}))

			areaChangePosition := container.New(layout.NewHBoxLayout(),
				widget.NewLabel("Change x,y of Mouse"),
				layout.NewSpacer(),
				widget.NewButton("Start", func() {
					content := container.New(layout.NewGridLayoutWithRows(3),
						layout.NewSpacer(),
						widget.NewLabel("Recording..."),
						layout.NewSpacer(),
					)
					newW.SetContent(content)
					newW.CenterOnScreen()
					newW.Resize(fyne.NewSize(300, 100))
					newW.Show()

					s := GetNewPositionOnScreen()

					c.StepParams = []interface{}{s.StepParams[0], s.StepParams[1]}
					newW.Close()

				}))

			generalContent := container.New(layout.NewVBoxLayout(),
				areaPreDelay,
				areaPostDelay,
			)
			configPositionContent := container.New(layout.NewVBoxLayout(),
				generalContent,
				areaChangePosition,
			)
			if c.StepAction == "MoveMouse" {
				newW.SetContent(configPositionContent)
			} else {
				newW.SetContent(generalContent)
			}
			newW.Resize(fyne.NewSize(100, 50))
			newW.CenterOnScreen()
			newW.Show()
		}),
	)
}

// GetWidget get test case widget
func (c *Step) GetWidget() *fyne.Container {
	return c.Widget
}
