package v4

import (
	"automateeverything.com/v2/maserver/routes"

	"path/filepath"
	"os"
	"log"

	"github.com/gin-gonic/gin"
)

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