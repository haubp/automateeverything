package v1

import (
	"automateeverything.com/v2/maserver/routes"

	"path/filepath"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Register ...
func Register(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "register_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "register_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "register_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "register_body.json")

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

// Discovered ...
func Discovered(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "discovered_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "discovered_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "discovered_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "discovered_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// FetchCustomUI ...
func FetchCustomUI(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "fetch_custom_ui_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "fetch_custom_ui_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "fetch_custom_ui_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "fetch_custom_ui_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// FetchSurvey ...
func FetchSurvey(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "fetch_survey_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "fetch_survey_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "fetch_survey_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "fetch_survey_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// FetchThreatHistory ...
func FetchThreatHistory(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "fetch_threat_history_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "fetch_threat_history_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "fetch_threat_history_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "fetch_threat_history_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// GetAccount ...
func GetAccount(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "get_account_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "get_account_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "get_account_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "get_account_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// GetConfigMetaScan ...
func GetConfigMetaScan(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "get_config_meta_scan_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "get_config_meta_scan_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "get_config_meta_scan_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "get_config_meta_scan_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// GetLicenseMetaScan ...
func GetLicenseMetaScan(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "get_license_meta_scan_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "get_license_meta_scan_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "get_license_meta_scan_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "get_license_meta_scan_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// GetPublicIP ...
func GetPublicIP(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "get_public_ip_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "get_public_ip_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "get_public_ip_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "get_public_ip_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

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

// LinkID ...
func LinkID(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "linkid_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "linkid_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "linkid_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "linkid_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// Pin ...
func Pin(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "pin_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "pin_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "pin_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "pin_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// PolicyCheck ...
func PolicyCheck(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "policy_check_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "policy_check_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "policy_check_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "policy_check_body.json")

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

// Result ...
func Result(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "result_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "result_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "result_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "result_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// FetchDisclaimer ...
func FetchDisclaimer(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "fetch_disclaimer_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "fetch_disclaimer_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "fetch_disclaimer_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "fetch_disclaimer_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// Unregister ...
func Unregister(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "unregister_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "unregister_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "unregister_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "unregister_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}