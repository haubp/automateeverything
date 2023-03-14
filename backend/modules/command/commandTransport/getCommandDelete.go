package commandTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/command/commandBiz"
	"backend_autotest/modules/command/commandModel"
	"backend_autotest/modules/command/commandStorage"
	"backend_autotest/modules/node/nodeBiz"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAndDeleteCommand(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data commandModel.Node

		var ok bool
		data.NodeId, ok = c.GetQuery("node_id")
		if !ok {
			panic(common.ErrInvalidRequest(nil))
		}

		store := commandStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := commandBiz.NewFindCommandBiz(store)

		result, err := biz.FindCommandAndDelete(c.Request.Context(), &data)
		pathAuto := fmt.Sprintf("./fileRepository/nodes/autoLog/%s", data.NodeId)
		pathAgent := fmt.Sprintf("./fileRepository/nodes/agentLog/%s", data.NodeId)
		nodeBiz.DeleteAutoLogFile(pathAuto)

		nodeBiz.DeleteAgentLogFile(pathAgent)
		if err != nil {
			c.JSON(400, err)
		}

		c.JSON(200, common.SimpleSuccessResponse(result))
	}
}
