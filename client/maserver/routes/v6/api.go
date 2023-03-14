package v6

import (
	"automateeverything.com/v2/maserver/routes"

	"path/filepath"
	"os"
	"log"

	"github.com/gin-gonic/gin"
)

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