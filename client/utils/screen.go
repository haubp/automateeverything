package utils

import (
	_ "image/png"

	"github.com/deluan/lookup"
	"github.com/go-vgo/robotgo"
)

func IsOnScreen(params []interface{}) bool {
	_isOnScreen := false

	sw, sh := robotgo.GetScreenSize()
  	screenBit := robotgo.CaptureScreen(0, 0, sw, sh)
  	defer robotgo.FreeBitmap(screenBit)
  	screenImg := robotgo.ToImage(screenBit)
	s := lookup.NewLookup(screenImg)

	templatePathList := []interface{}{params[0].(string)}

	template := LoadImageFromFile(templatePathList)

	pp, _ := s.FindAll(template, 0.9)

	if len(pp) > 0 {
		_isOnScreen = true;
	}

	return _isOnScreen;
}