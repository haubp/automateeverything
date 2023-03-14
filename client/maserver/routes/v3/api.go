package v3

import (
	"automateeverything.com/v2/maserver/routes"

	"path/filepath"
	"os"
	"log"

	"github.com/gin-gonic/gin"
)

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

// ScanResult ...
func ScanResult(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "scan_result_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "scan_result_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "scan_result_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "scan_result_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}

// GetCertManifest ...
func GetCertManifest(c *gin.Context) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	reqHeaderPath := filepath.Join(pwd, "maserver-live", "request", "get_cert_manifest_header.json")
	reqBodyPath := filepath.Join(pwd, "maserver-live", "request", "get_cert_manifest_body.json")
	resHeaderPath := filepath.Join(pwd, "maserver-live", "response", "get_cert_manifest_header.json")
	resBodyPath := filepath.Join(pwd, "maserver-live", "response", "get_cert_manifest_body.json")

	routes.DefaultReadHandler(c, reqHeaderPath, reqBodyPath)
	routes.DefaultWriteHandler(c, resHeaderPath, resBodyPath)
}