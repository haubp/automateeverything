package nodeTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeBiz"
	"backend_autotest/modules/node/nodeModel"
	"backend_autotest/modules/node/nodeStorage"
	"github.com/gin-gonic/gin"
)

func NodePostResult(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data nodeModel.Node
		var ok bool
		data.NodeId, ok = c.GetQuery("node_id")
		if !ok {
			panic(common.ErrInvalidRequest(nil))
		}

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := nodeStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := nodeBiz.NewAddResultBiz(store)
		_, err := biz.AddNewResult(c.Request.Context(), &data)
		if err != nil {
			c.JSON(400, err)
			return
		}

		c.JSON(200, common.SimpleSuccessResponse("success"))
	}
}
