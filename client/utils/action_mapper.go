package utils

func ActionsMap() map[string]interface{} {
	actionsMap := map[string]interface{} {
		"IsOnScreen": IsOnScreen,
		"MoveMouse": MoveMouse,
		"LeftClick": LeftClick,
		"RightClick": RightClick,
		"ScrollMouse": ScrollMouse,
		"CheckLog": CheckLog,
	}

	return actionsMap
}