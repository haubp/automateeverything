package utils

import (
	"github.com/go-vgo/robotgo"
  )

// KeyboardTyping typing 1: string to type
func KeyboardTyping(params []interface{}) bool {
	str := params[0].(string)
	robotgo.TypeStr(str)
	return true
}