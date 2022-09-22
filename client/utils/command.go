package utils

import (
	"log"
	"os/exec"
)

// RunCommand run command 0: command to run
func RunCommand(params []interface{}) bool {
	cmdStr := params[0].(string)
	out, err := exec.Command(cmdStr).Output()
	if err != nil {
		return false
	}
	log.Println("Command stdout:\n", string(out))
	return true
}
