package template

import (
	"fmt"
	hook "github.com/robotn/gohook"
	"github.com/go-vgo/robotgo"
)

type Step struct {
	StepName string `json:"step_name"`
	StepAction string `json:"step_action"`
	StepParams []interface{} `json:"step_params"`
	PreSleep int `json:"step_pre_sleep"`
	PostSleep int `json:"step_post_sleep"`
}

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
      	StepName: "any", 
      	StepAction: "MoveMouse", 
      	StepParams: []interface{}{ x, y }, 
      	PreSleep: 1000, 
      	PostSleep: 1000})
	steps = append(steps, Step{
    	StepName: "any", 
    	StepAction: "LeftClick", 
    	StepParams: []interface{}{ false }, 
    	PreSleep: 1000, 
    	PostSleep: 1000})
  })

  hook.Register(hook.KeyDown, []string{"r"}, func(e hook.Event) {
	x, y := robotgo.GetMousePos()
    fmt.Println("right click", x, y)
  	steps = append(steps, Step{
    	StepName: "any", 
    	StepAction: "MoveMouse", 
    	StepParams: []interface{}{ x, y }, 
    	PreSleep: 1000, 
    	PostSleep: 1000})
	steps = append(steps, Step{
    	StepName: "any", 
    	StepAction: "RightClick", 
    	StepParams: []interface{}{ false }, 
    	PreSleep: 1000, 
    	PostSleep: 1000})
  })

  hook.Register(hook.KeyDown, []string{"s"}, func(e hook.Event) {
    x, y := robotgo.GetMousePos()
      fmt.Println("scroll", x, y)
    steps = append(steps, Step{
      StepName: "any", 
      StepAction: "MoveMouse", 
      StepParams: []interface{}{ x, y }, 
      PreSleep: 1000, 
      PostSleep: 1000})
    steps = append(steps, Step{
      StepName: "any", 
      StepAction: "ScrollMouse", 
      StepParams: []interface{}{ 100 }, 
      PreSleep: 1000, 
      PostSleep: 1000})
    })

	hook.Register(hook.KeyDown, []string{"t"}, func(e hook.Event) {
		steps = append(steps, Step{
		  StepName: "any", 
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

  fmt.Println(steps)

  return steps;
}