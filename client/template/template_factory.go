package template

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
	for _, group := range t.TestCategoryGroups {
		totalOfTestCases += float64(len(group.TestGroupTestCases))
	}
	percentOfEachTestCase := 100 / totalOfTestCases
	for _, group := range t.TestCategoryGroups {
		resultTelevision.WriteAndExpand("	Run test for group " + group.TestGroupName)
		for _, test := range group.TestGroupTestCases {
			resultTelevision.WriteAndExpand("		Run test for test " + test.TestCaseName)
			test.Result = "Success"
			var capturedTime time.Time
			for _, step := range test.TestCaseSteps {
				resultTelevision.WriteAndExpand("			Run step " + step.StepName)
				if step.StepAction == "CaptureTime" {
					capturedTime = time.Now()
					resultTelevision.WriteAndExpand("			Capture time")
				} else {
					if step.StepAction == "CheckLog" {
						step.StepParams = append(step.StepParams, capturedTime)
					}
					time.Sleep(time.Duration(step.PreSleep) * time.Millisecond)
					if actionsMap[step.StepAction].(func([]interface{}) bool)(step.StepParams) {
						resultTelevision.WriteAndExpand("				" + "Success")
					} else {
						resultTelevision.WriteAndExpand("				" + "Failed")
						test.Result = "Failed at step " + step.StepName
						break
					}
					time.Sleep(time.Duration(step.PostSleep) * time.Millisecond)
				}
				resultContainer.ScrollToBottom()
			}
			progressBar.SetValue(progressBar.Value + percentOfEachTestCase)
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
