package nodeTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeModel"
	"fmt"
	"github.com/gin-gonic/gin"
)

func PostAutoLogFile(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data nodeModel.Node
		var ok bool
		data.NodeId, ok = c.GetQuery("node_id")
		if !ok {
			panic(common.ErrInvalidRequest(nil))
		}

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.SaveUploadedFile(fileHeader, fmt.Sprintf("./fileRepository/nodes/autoLog/%s", data.NodeId))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(200, common.SimpleSuccessResponse("success"))
	}
}
