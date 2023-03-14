package v2

import (
	"automateeverything.com/v2/maserver/routes"

	"path/filepath"
	"os"
	"log"

	"github.com/gin-gonic/gin"
)

// FetchPrivacy ...
func FetchPrivacy(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "fetch_privacy_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "fetch_privacy_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "fetch_privacy_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "fetch_privacy_body.json")

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