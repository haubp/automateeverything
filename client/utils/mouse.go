package utils

import (
	"github.com/go-vgo/robotgo"
)

// GetCurrentMousePosition get current mouse position
func GetCurrentMousePosition(params []interface{}) (mx, my int) {
	mx, my = robotgo.GetMousePos()
	return
}

// MoveMouse move mouse
func MoveMouse(params []interface{}) bool {
	robotgo.MouseSleep = 100
	x := params[0].(float64)
	y := params[1].(float64)
	robotgo.Move(int(x), int(y))
	return true
}

// LeftClick left click
func LeftClick(params []interface{}) bool {
	robotgo.Click("left", params[0].(bool))
	return true
}

// RightClick right click
func RightClick(params []interface{}) bool {
	robotgo.Click("right", params[0].(bool))
	return true
}

// ScrollMouse scroll mouse
func ScrollMouse(params []interface{}) bool {
	robotgo.MouseSleep = 100
	x := int(params[0].(float64))
	robotgo.ScrollMouse(x, "down")
	return true
}

// DragMouse drag mouse
func DragMouse(params []interface{}) bool {
	robotgo.MouseSleep = 100
	x := params[0].(float64)
	y := params[1].(float64)
	robotgo.DragSmooth(int(x), int(y))
	return true
}

// MouseSleep sleep
func MouseSleep(params []interface{}) bool {
	robotgo.MilliSleep(int(params[0].(float64)))
	return true
}