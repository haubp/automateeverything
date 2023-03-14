package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"log"
	"bytes"
	
	"github.com/gin-gonic/gin"
)

func readJSONFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Read json file at %s with err: %s\n", path, err.Error())
		return nil
	}
	if !json.Valid(data) {
		fmt.Printf("Invalid json format\n")
	}

	return data
}

// DefaultWriteHandler ...
func DefaultWriteHandler(c *gin.Context, resHeaderPath string, resBodyPath string) {
	// Serving Reponse
	byteResp := readJSONFile(resBodyPath)
	c.Data(http.StatusOK, "application/json", byteResp)
}

func createKeyValuePairs(m map[string][]string) string {
    b := new(bytes.Buffer)
    for key, value := range m {
        fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
    }
    return b.String()
}

// DefaultReadHandler ...
func DefaultReadHandler(c *gin.Context, reqHeaderPath string, reqBodyPath string) {
	// Save request header to file
	err := ioutil.WriteFile( reqHeaderPath, []byte(createKeyValuePairs(c.Request.Header)), 0644)

	// Save request response to file
	bodyJSONData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile( reqBodyPath, bodyJSONData, 0644)
	if err != nil {
		log.Println(err)
	}
}
