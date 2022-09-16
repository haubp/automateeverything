package runner

import (
	"fmt"
	"time"

	"automateeverything.com/v1/template"
	"automateeverything.com/v1/utils"
)

func ExecuteTest(templatePath string) {
	actionsMap := utils.ActionsMap()

	testTemplate, err := template.CreateTestFromJsonFile(templatePath)
	if err != nil {
		fmt.Println("Error ", err.Error())
	}

	fmt.Println("Run test for category", testTemplate.TestCategoryName)
	for _, group := range testTemplate.TestCategoryGroups {
		fmt.Println("	Run test for group", group.TestGroupName)
		for _, test := range group.TestGroupTests {
			fmt.Println("		Run test for test", test.TestName)
			capturedTime := time.Now()
			for _, step := range test.TestSteps {
				fmt.Println("			Run step", step.StepName)
				if step.StepAction == "CaptureTime" {
					capturedTime = time.Now()
				} else {
					if step.StepAction == "CheckLog" {
						step.StepParams = append(step.StepParams, capturedTime)
					}

					time.Sleep(time.Duration(step.PreSleep) * time.Millisecond)
					if actionsMap[step.StepAction].(func([]interface{}) bool)(step.StepParams) {
						fmt.Println("				", "Success")
					} else {
						fmt.Println("				", "Failed")
						break
					}
					time.Sleep(time.Duration(step.PostSleep) * time.Millisecond)
				}
			}
		}
	}
}
