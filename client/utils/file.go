package utils

import (
	"image"
	"time"
	"regexp"
	"bufio"
    "strings"
    "log"
    "os"
)

// LoadImageFromFile 0: image path
func LoadImageFromFile(params []interface{}) image.Image {
	imageFile, _ := os.Open(params[0].(string))
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}

// CheckLog 0: log path 1: log pattern 2: start at
func CheckLog(params []interface{}) bool {
	timeLogLayout := "01-02-06 (15:04:05)"

	logPath := params[0].(string)
	logPattern := params[1].(string)
	startTime := params[2].(time.Time)

	file, err := os.Open(logPath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "<?>") {
			line = decryptLog(line)
		}

		r, _ := regexp.Compile(`^\[(.*)\] ([\d-]{8} \(.*\)) GMT\-\d{2} \(\d+\) (.*)$`)

		matches := r.FindAllStringSubmatch(line, -1)
		if matches != nil {
			// logType := matches[0][1]
			logTime := matches[0][2]
			logText := matches[0][3]

			tCurrent, err := time.Parse(timeLogLayout, logTime)

			if err != nil {
				panic(err)
			}
			
			if startTime.Before(tCurrent) {
				if strings.Contains(logText, logPattern) {
					return true
				}
			} 
		}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return false
}
