package utils

import (
	"github.com/go-vgo/robotgo"
)

func GetCurrentMousePosition(params []interface{}) (mx, my int) {
	mx, my = robotgo.GetMousePos()
	return
}

func MoveMouse(params []interface{}) bool {
	robotgo.MouseSleep = 100
	x := params[0].(float64)
	y := params[1].(float64)
	robotgo.Move(int(x), int(y))
	return true
}

func LeftClick(params []interface{}) bool {
	robotgo.Click("left", params[0].(bool))
	return true
}

func RightClick(params []interface{}) bool {
	robotgo.Click("right", params[0].(bool))
	return true
}

func ScrollMouse(params []interface{}) bool {
	robotgo.MouseSleep = 100
	x := int(params[0].(float64))
	robotgo.ScrollMouse(x, "down")
	return true
}

func DragMouse(params []interface{}) bool {
	robotgo.MouseSleep = 100
	x := params[0].(float64)
	y := params[1].(float64)
	robotgo.DragSmooth(int(x), int(y))
	return true
}

func MouseSleep(params []interface{}) bool {
	robotgo.MilliSleep(int(params[0].(float64)))
	return true
}