package maserver

import (
	v1 "automateeverything.com/v2/maserver/routes/v1"
	v2 "automateeverything.com/v2/maserver/routes/v2"
	v3 "automateeverything.com/v2/maserver/routes/v3"
	v4 "automateeverything.com/v2/maserver/routes/v4"
	v5 "automateeverything.com/v2/maserver/routes/v5"
	v6 "automateeverything.com/v2/maserver/routes/v6"

	"github.com/gin-gonic/gin"
)

// InitMaServer ...
func InitMaServer() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/gears/api/1/agent")
	{
		apiV1.POST("/register", v1.Register)
		apiV1.POST("/fetch_config", v1.FetchConfig)
		apiV1.POST("/discovered", v1.Discovered)
		apiV1.POST("/fetch_custom_ui", v1.FetchCustomUI)
		apiV1.POST("/fetch_survey", v1.FetchSurvey)
		apiV1.POST("/fetch_threat_history", v1.FetchThreatHistory)
		apiV1.POST("/get_account", v1.GetAccount)
		apiV1.POST("/get_config/metascan", v1.GetConfigMetaScan)
		apiV1.POST("/get_license/metascan", v1.GetLicenseMetaScan)
		apiV1.POST("/get_public_ip", v1.GetPublicIP)
		apiV1.POST("/imhere", v1.ImHere)
		apiV1.POST("/linkid", v1.LinkID)
		apiV1.POST("/pin", v1.Pin)
		apiV1.POST("/policy_check", v1.PolicyCheck)
		apiV1.POST("/report_soh", v1.ReportSOH)
		apiV1.POST("/report_sot", v1.ReportSOT)
		apiV1.POST("/report_sov", v1.ReportSOV)
		apiV1.POST("/result", v1.Result)
		apiV1.POST("/fetch_disclaimer", v1.FetchDisclaimer)
		apiV1.POST("/unregister", v1.Unregister)
	}

	apiV2 := router.Group("/gears/api/2/agent")
	{
		apiV2.POST("/fetch_privacy", v2.FetchPrivacy)
		apiV2.POST("/report_sot", v2.ReportSOT)
		apiV2.POST("/report_sov", v2.ReportSOV)
	}

	apiV3 := router.Group("/gear/api/3/agent")
	{
		apiV3.POST("/discovered", v3.Discovered)
		apiV3.POST("/fetch_config", v3.FetchConfig)
		apiV3.POST("/get_account", v3.GetAccount)
		apiV3.POST("/imhere", v3.ImHere)
		apiV3.POST("/linked", v3.LinkID)
		apiV3.POST("/pin", v3.Pin)
		apiV3.POST("/register", v3.Register)
		apiV3.POST("/report_soh", v3.ReportSOH)
		apiV3.POST("/report_sot", v3.ReportSOT)
		apiV3.POST("/result", v3.Result)
		apiV3.POST("/unregister",  v3.Unregister)
		apiV3.POST("/scan_result", v3.ScanResult)
		apiV3.POST("/get_cert_manifest", v3.GetCertManifest)
	}

	apiV4 := router.Group("/gear/api/4/agent")
	{
		apiV4.POST("/report_sot", v4.ReportSOT)
	}

	apiV5 := router.Group("/gears/api/5/agent")
	{
		apiV5.POST("/imhere", v5.ImHere)
		apiV5.POST("/report_sov", v5.ReportSOV)
		apiV5.POST("/report_soh", v5.ReportSOH)
		apiV5.POST("/report_sot", v5.ReportSOT)
		apiV5.POST("/check_message", v5.CheckMessage)
		apiV5.POST("/fetch_config", v5.FetchConfig)
		apiV5.POST("/fetch_key", v5.FetchKey)
		apiV5.POST("/feedback", v5.Feedback)
		apiV5.POST("/smart_scan", v5.SmartScan)
		apiV5.POST("/assign_group", v5.AssignGroup)
		apiV5.POST("/status", v5.Status)
		apiV5.POST("/fetch_report_rules", v5.FetchReportRules)
		apiV5.POST("/compliance_status", v5.ComplianceStatus)
		apiV5.POST("/fetch_workflows", v5.FetchWorkflows)
		apiV5.POST("/action_result", v5.ActionResult)
		apiV5.POST("/fetch_processed_data", v5.FetchProcessedData)
		apiV5.POST("/remediation_status", v5.RemediationStatus)
		apiV5.GET("/download_custom_script", v5.DownloadCustomScript)
	}

	apiV6 := router.Group("/gears/api/6/agent")
	{
		apiV6.POST("/report_soh", v6.ReportSOH)
	}

	return router
}
