package utils

// ActionsMap action mapper
func ActionsMap() map[string]interface{} {
	actionsMap := map[string]interface{} {
		"IsOnScreen": IsOnScreen,
		"MoveMouse": MoveMouse,
		"LeftClick": LeftClick,
		"RightClick": RightClick,
		"ScrollMouse": ScrollMouse,
		"CheckLog": CheckLog,
		"KeyboardTyping": KeyboardTyping,
		"RunCommand": RunCommand,
		"CheckClientRequestToMA": CheckClientRequestToMA,
	}
	return actionsMap
}
