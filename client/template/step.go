package template

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"github.com/go-vgo/robotgo"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

// Step type
type Step struct {
	StepName string `json:"step_name"`
	StepAction string `json:"step_action"`
	StepParams []interface{} `json:"step_params"`
	PreSleep int `json:"step_pre_sleep"`
	PostSleep int `json:"step_post_sleep"`
	Widget *fyne.Container `json:"-"`
	W fyne.Window `json:"-"`
	A fyne.App `json:"-"`
}

// MakeSteps auto
func MakeSteps() []Step {
  steps := []Step{}

  fmt.Println("Start Hook")

  hook.Register(hook.KeyDown, []string{"d"}, func(e hook.Event) {
	x, y := robotgo.GetMousePos()
    fmt.Println("double click", x, y)
  	steps = append(steps, Step{
    	StepName: "any", 
    	StepAction: "MoveMouse", 
    	StepParams: []interface{}{ x, y }, 
    	PreSleep: 1000, 
    	PostSleep: 1000})
	steps = append(steps, Step{
    	StepName: "any", 
    	StepAction: "LeftClick", 
    	StepParams: []interface{}{ true }, 
    	PreSleep: 1000, 
    	PostSleep: 1000})
  })

  hook.Register(hook.KeyDown, []string{"c"}, func(e hook.Event) {
	x, y := robotgo.GetMousePos()
    fmt.Println("click", x, y)
    steps = append(steps, Step{
      	StepName: "Move Mouse", 
      	StepAction: "MoveMouse", 
      	StepParams: []interface{}{ x, y }, 
      	PreSleep: 1000, 
      	PostSleep: 1000})
	steps = append(steps, Step{
    	StepName: "Left Click", 
    	StepAction: "LeftClick", 
    	StepParams: []interface{}{ false }, 
    	PreSleep: 1000, 
    	PostSleep: 1000})
  })

  hook.Register(hook.KeyDown, []string{"r"}, func(e hook.Event) {
	x, y := robotgo.GetMousePos()
    fmt.Println("right click", x, y)
  	steps = append(steps, Step{
    	StepName: "Move Mouse", 
    	StepAction: "MoveMouse", 
    	StepParams: []interface{}{ x, y }, 
    	PreSleep: 1000, 
    	PostSleep: 1000})
	steps = append(steps, Step{
    	StepName: "Right Click", 
    	StepAction: "RightClick", 
    	StepParams: []interface{}{ false }, 
    	PreSleep: 1000, 
    	PostSleep: 1000})
  })

  hook.Register(hook.KeyDown, []string{"s"}, func(e hook.Event) {
    x, y := robotgo.GetMousePos()
      fmt.Println("scroll", x, y)
    steps = append(steps, Step{
      StepName: "Move Mouse", 
      StepAction: "MoveMouse", 
      StepParams: []interface{}{ x, y }, 
      PreSleep: 1000, 
      PostSleep: 1000})
    steps = append(steps, Step{
      StepName: "Scroll Mouse", 
      StepAction: "ScrollMouse", 
      StepParams: []interface{}{ 100 }, 
      PreSleep: 1000, 
      PostSleep: 1000})
    })

	hook.Register(hook.KeyDown, []string{"t"}, func(e hook.Event) {
		steps = append(steps, Step{
		  StepName: "Capture Time", 
		  StepAction: "CaptureTime", 
		  StepParams: []interface{}{ }, 
		  PreSleep: 100, 
		  PostSleep: 100})
	})

  hook.Register(hook.KeyDown, []string{"q"}, func(e hook.Event) {
    fmt.Println("exit")
	hook.End()
  })

  fmt.Println("Listening")

  s := hook.Start()
  <-hook.Process(s)

  return steps;
}

// InitContext init step context
func (c * Step) InitContext(a fyne.App, w fyne.Window) {
	c.W = w
	c.A = a
	c.Widget = container.New(layout.NewHBoxLayout(),
		canvas.NewText(c.StepName, color.RGBA{0xD8, 0xD8, 0xD8, 1}), 
		layout.NewSpacer(),
		widget.NewButton("X", func(){
		
		}),
	)
}

// GetWidget get test case widget
func (c * Step) GetWidget() *fyne.Container {
	return c.Widget
}