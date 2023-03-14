package main

import (
	"automateeverything.com/v2/utils"
	"fmt"
	"log"
	"os"
	"time"
)

func testFolderExit() {
	log.SetOutput(os.Stdout)
	params := []interface{}{
		"IF EXIST \"C:\\Program Files (x86)\\OPSWAT\\GEARS Client\" ECHO exists",
		"exists",
	}
	fmt.Println(utils.RunCommand(params))
}

func testRegisterNothing() {
	params := []interface{}{
		"register",
		"agent_type",
		"0",
	}
	fmt.Println(utils.CheckClientRequestToMA(params))
}

func testRemoveFolder() {
	log.SetOutput(os.Stdout)
	params := []interface{}{
		"powershell.exe -File \"C:\\Users\\hau.bui\\OneDrive - OPSWAT\\Desktop\\automation\\bit-auto\\automation\\client\\RT\\script\\cleanup.ps1\"",
		"Finish",
	}
	fmt.Println(utils.RunCommand(params))
}

func testCheckLog() {
	var capturedTime time.Time
	params := []interface{}{
		"..\\GEARS-Agent.log",
		"KeepAlive",
	}
	params = append(params, capturedTime)

	fmt.Println(utils.CheckLog(params))
}

func main() {
	testCheckLog()
}
