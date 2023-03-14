package template

import ()

// MAResponse ...
type MAResponse struct {
	RequestPath string 		`json:"request_path"`
	ResponseHeader string 	`json:"response_header"`
	ResponseBody string 	`json:"response_body"`
}
