package v5

import (
	"automateeverything.com/v2/maserver/routes"

	"path/filepath"
	"os"
	"log"

	"github.com/gin-gonic/gin"
)

// ImHere ...
func ImHere(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "imhere_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "imhere_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "imhere_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "imhere_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// ReportSOV ...
func ReportSOV(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "report_sov_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "report_sov_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "report_sov_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "report_sov_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// ReportSOH ...
func ReportSOH(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "report_soh_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "report_soh_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "report_soh_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "report_soh_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// ReportSOT ...
func ReportSOT(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "report_sot_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "report_sot_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "report_sot_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "report_sot_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// CheckMessage ...
func CheckMessage(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "check_message_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "check_message_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "check_message_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "check_message_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// FetchConfig ...
func FetchConfig(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "fetch_config_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "fetch_config_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "fetch_config_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "fetch_config_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// FetchKey ...
func FetchKey(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "fetch_key_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "fetch_key_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "fetch_key_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "fetch_key_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// Feedback ...
func Feedback(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "feedback_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "feedback_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "feedback_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "feedback_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// SmartScan ...
func SmartScan(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "smart_scan_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "smart_scan_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "smart_scan_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "smart_scan_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// AssignGroup ...
func AssignGroup(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "assign_group_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "assign_group_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "assign_group_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "assign_group_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// Status ...
func Status(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "status_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "status_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "status_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "status_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// FetchReportRules ...
func FetchReportRules(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "fetch_report_rules_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "fetch_report_rules_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "fetch_report_rules_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "fetch_report_rules_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// ComplianceStatus ...
func ComplianceStatus(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "compliance_status_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "compliance_status_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "compliance_status_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "compliance_status_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// FetchWorkflows ...
func FetchWorkflows(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "fetch_workflows_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "fetch_workflows_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "fetch_workflows_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "fetch_workflows_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// ActionResult ...
func ActionResult(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "action_result_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "action_result_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "action_result_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "action_result_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// FetchProcessedData ...
func FetchProcessedData(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "fetch_processed_data_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "fetch_processed_data_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "fetch_processed_data_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "fetch_processed_data_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// RemediationStatus ...
func RemediationStatus(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "remediation_status_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "remediation_status_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "remediation_status_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "remediation_status_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// DownloadCustomScript ...
func DownloadCustomScript(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "download_custom_script_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "download_custom_script_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "download_custom_script_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "download_custom_script_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}