package template

import (
	"encoding/json"
	"fmt"
	cp "github.com/otiai10/copy"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"automateeverything.com/v2/custom"
	"automateeverything.com/v2/utils"
)

// SelectedTestCase Global Selected Test Case => TODO: Need to remove global variable
var SelectedTestCase *TestCase

// CreateTestFromJSONFile Creat Test Struct from json file
func CreateTestFromJSONFile(templatePath string) (TestCategory, error) {
	templateFile, err := os.Open(templatePath)
	if err != nil {
		log.Println(err)
		return TestCategory{}, err
	}
	log.Println("Successfully Opened ", templatePath)
	defer templateFile.Close()

	templateByteValue, _ := ioutil.ReadAll(templateFile)

	var testTemplate TestCategory
	json.Unmarshal(templateByteValue, &testTemplate)

	return testTemplate, nil
}

// CreateTestFromBytes Creat Test Struct from json file
func CreateTestFromBytes(b []byte) (TestCategory, error) {
	var testTemplate TestCategory
	json.Unmarshal(b, &testTemplate)

	return testTemplate, nil
}

// CreateJSONFileFromTemplate Creat json file from Test Struct
func CreateJSONFileFromTemplate(templatePath string, test TestCategory) {
	file, e := json.MarshalIndent(test, "", " ")
	if e != nil {
		log.Println(e)
	}
	e = ioutil.WriteFile(templatePath, file, 0644)
	if e != nil {
		log.Println(e)
	}
}

// ExecuteTestFromTemplate execute test from file bytes
func ExecuteTestFromTemplate(t *TestCategory, resultTelevision *custom.FixSizeLabel, progressBar *widget.ProgressBar, resultContainer *container.Scroll, reportNameEntry *widget.Entry) {
	actionsMap := utils.ActionsMap()
	resultTelevision.WriteAndExpand("Run test for category " + t.TestCategoryName)
	totalOfTestCases := 0.0
	resultTelevision.WriteAndExpand("		Clear MA Request data")
	clearRequestFolder()
	resultTelevision.WriteAndExpand("		Clear MA Request data Done")
	for _, group := range t.TestCategoryGroups {
		totalOfTestCases += float64(len(group.TestGroupTestCases))
	}
	percentOfEachTestCase := 100 / totalOfTestCases

	for _, group := range t.TestCategoryGroups {
		resultTelevision.WriteAndExpand("	Run test for group " + group.TestGroupName)
		for _, test := range group.TestGroupTestCases { // bk
			RollbackResponse("./maserver-live/response_backup", "./maserver-live/response")

			resultTelevision.WriteAndExpand("		Run test for test " + test.TestCaseName)
			initMAServer(test)

			test.Result = "Success"
			var capturedTime time.Time
			for _, step := range test.TestCaseSteps {
				resultTelevision.WriteAndExpand("			Run step " + step.StepName)
				resultContainer.ScrollToBottom()
				if step.StepAction == "CaptureTime" {
					capturedTime = time.Now()
					resultTelevision.WriteAndExpand("			Capture time")
				} else {
					if step.StepAction == "CheckLog" {
						step.StepParams = append(step.StepParams, capturedTime)
					}
					resultTelevision.WriteAndExpand("			Pre sleep " + strconv.Itoa(step.PreSleep/1000) + " second")
					time.Sleep(time.Duration(step.PreSleep) * time.Millisecond)
					if actionsMap[step.StepAction].(func([]interface{}) bool)(step.StepParams) {
						resultTelevision.WriteAndExpand("				" + "Success")
					} else {
						resultTelevision.WriteAndExpand("				" + "Failed")
						test.Result = "Failed at step " + step.StepName + "-" + time.Now().String()
						break
					}
					resultTelevision.WriteAndExpand("			Post sleep " + strconv.Itoa(step.PostSleep/1000) + " second")

					time.Sleep(time.Duration(step.PostSleep) * time.Millisecond)
				}
				resultContainer.ScrollToBottom()
			}
			progressBar.SetValue(progressBar.Value + percentOfEachTestCase)

			// bk
			RollbackResponse("./maserver-live/response_backup", "./maserver-live/response")
		}
	}
	resultTelevision.WriteAndExpand("Report")
	if reportNameEntry.Text != "" {
		CreateJSONFileFromTemplate(reportNameEntry.Text, *t)
	} else {
		CreateJSONFileFromTemplate("report.json", *t)
	}

	resultTelevision.WriteAndExpand("Finish")
	resultContainer.ScrollToBottom()
}

func RollbackResponse(source string, des string) {

	if err := cp.Copy(source, des); err != nil {
		log.Println(err)
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func clearRequestFolder() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	items, _ := ioutil.ReadDir(filepath.Join(pwd, "maserver-live", "request"))
	for _, item := range items {
		if !item.IsDir() {
			err := os.WriteFile(filepath.Join(pwd, "maserver-live", "request", item.Name()), []byte{}, 0644)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func initMAServer(tc *TestCase) {

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	for _, response := range tc.MAResponseL {

		file_name := fmt.Sprint("./maserver-live/response/", fmt.Sprint(response.RequestPath, "_header.json"))
		if exists(file_name) != true {
			log.Println("File not found: ", file_name)
			continue
		}

		file_name = fmt.Sprint("./maserver-live/response/", fmt.Sprint(response.RequestPath, "_body.json"))
		if exists(file_name) != true {
			log.Println("File not found: ", file_name)
			continue
		}

		err := os.WriteFile(filepath.Join(pwd, "maserver-live", "response", fmt.Sprint(response.RequestPath, "_header.json")), []byte(response.ResponseHeader), 0644)
		if err != nil {
			log.Println(err)
		}
		err = os.WriteFile(filepath.Join(pwd, "maserver-live", "response", fmt.Sprint(response.RequestPath, "_body.json")), []byte(response.ResponseBody), 0644)
		if err != nil {
			log.Println(err)
		}
	}
}
