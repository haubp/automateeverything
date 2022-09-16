package template

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

func CreateTestFromJsonFile(templatePath string) (TestCategory, error) {
	templateFile, err := os.Open(templatePath)
	if err != nil {
    	fmt.Println(err)
		return TestCategory{}, err
	}
	fmt.Println("Successfully Opened ", templatePath)
	defer templateFile.Close()

    templateByteValue, _ := ioutil.ReadAll(templateFile)

	var testTemplate TestCategory
	json.Unmarshal(templateByteValue, &testTemplate)

	return testTemplate, nil
}

func CreateJsonFileFromTemplate(templatePath string, test TestCategory) {
	file, _ := json.MarshalIndent(test, "", " ")
	_ = ioutil.WriteFile(templatePath, file, 0644)
}