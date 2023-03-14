package nodeTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func GetTemplate(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data nodeModel.Node
		var ok bool
		data.UserName, ok = c.GetQuery("user_name")
		if !ok {
			panic(common.ErrInvalidRequest(nil))
		}

		var file, err = os.OpenFile(fmt.Sprintf("./fileRepository/users/%s", data.UserName), os.O_RDWR, 0644)

		if err != nil {
			panic(err)
		}

		defer file.Close()

		bytefile, err := ioutil.ReadAll(file)

		if err != nil {
			panic(err)
		}

		c.JSON(200, bytefile)
	}
}
