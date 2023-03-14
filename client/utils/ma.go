package utils

import (
	"os"
	"path/filepath"
	"log"
	"fmt"
	"encoding/json"
	"strings"
	"strconv"
)

// CheckClientRequestToMA param: requestPath, queryEntry, valueToCompare
func CheckClientRequestToMA(params []interface{}) bool {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	requestPath := params[0].(string)
	requestQueryEntry := params[1].(string)
	valueToCompare := params[2].(string)

	var bodyData []byte

	switch(requestPath) {
	case "register":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "register_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "unregister":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "unregister_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "imhere":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "imhere_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "action_result":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "action_result_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "assign_group":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "assign_group_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "check_message":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "check_message_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "compliance_status":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "compliance_status_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "discovered":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "discovered_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "download_custom_script":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "download_custom_script_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "feedback":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "feedback_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "fetch_config":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "fetch_config_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "fetch_custom_ui":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "fetch_custom_ui_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "fetch_disclaimer":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "fetch_disclaimer_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "fetch_key":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "fetch_key_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "fetch_privacy":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "fetch_privacy_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "fetch_processed_data":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "fetch_processed_data_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "fetch_report_rules":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "fetch_report_rules_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "fetch_survey":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "fetch_survey_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "fetch_thread_history":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "fetch_threat_history_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "fetch_workflows":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "fetch_workflows_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "get_account":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "get_account_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "get_cert_manifest":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "get_cert_manifest_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "get_config_meta_scan":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "get_config_meta_scan_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "get_license_meta_scan":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "get_license_meta_scan_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "get_public_ip":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "get_public_ip_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "linkid":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "linkid_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "pin":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "pin_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "policy_check":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "policy_check_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "remediation_status":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "remediation_status_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "report_soh":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "report_soh_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "report_sot":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "report_sot.json"))
		if err != nil {
			log.Println(err)
		}
	case "report_sov":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "report_sov_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "result":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "result_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "scan_result":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "scan_result_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "smart_scan":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "smart_scan_body.json"))
		if err != nil {
			log.Println(err)
		}
	case "status":
		bodyData, err = os.ReadFile(filepath.Join(pwd, "maserver-live", "request", "status_body.json"))
		if err != nil {
			log.Println(err)
		}
	}

	bodyMap := make(map[string]interface{})
	json.Unmarshal(bodyData, &bodyMap)

	if len(bodyMap) == 0 && valueToCompare == "" {
		return true
	}

	if len(bodyMap) == 0 {
		log.Println("Request empty")
		return false
	}

	requestQueryEntryList := strings.Split(requestQueryEntry, ".")
	fmt.Println(requestQueryEntryList)

	for i := 0; i < len(requestQueryEntryList) - 1; i++ {
		bodyMap = bodyMap[requestQueryEntryList[i]].(map[string]interface{})
		if bodyMap == nil {
			log.Println("Could not query entry")
			return false
		}
	}

	switch bodyMap[requestQueryEntryList[len(requestQueryEntryList) - 1]].(type) {
	case string:
		if bodyMap[requestQueryEntryList[len(requestQueryEntryList) - 1]].(string) != valueToCompare {
			return false
		}
	case float64:
		if strconv.Itoa(int(bodyMap[requestQueryEntryList[len(requestQueryEntryList) - 1]].(float64))) != valueToCompare {
			return false
		}
	}

	return true
}
