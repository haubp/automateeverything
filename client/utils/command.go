package utils

import (
	"log"
	"os/exec"
	"strings"
	"io"
)

// RunCommand run command 0: command to run
func RunCommand(params []interface{}) bool {
	cmdStr := params[0].(string)
	expectedStr := params[1].(string)

	cmd := exec.Command("cmd")
	stdin, err := cmd.StdinPipe()
	if err != nil {
	    log.Fatal(err)
	}
	go func() {
	    defer stdin.Close()
	    io.WriteString(stdin, cmdStr + "\n")
	}()
	s, err := cmd.Output()
	if err != nil {
	    log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outputString := string(s)

	log.Println("Run command result:\n", outputString)

	sList := strings.Split(outputString, "\n")
	index := -1
	for i, line := range sList {
		if strings.Contains(line, cmdStr) {
			index = i + 1
		}
	}
	if index == -1 {
		return false
	}
	for ; index <  len(sList); index++ {
		if strings.Contains(sList[index], expectedStr) {
			return true
		}
	}
	return false
}
