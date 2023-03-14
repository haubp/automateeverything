package nodeTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeBiz"
	"backend_autotest/modules/node/nodeStorage"
	"github.com/gin-gonic/gin"
)

func NodeGetResult(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data string
		var ok bool
		data, ok = c.GetQuery("node_id")
		if !ok {
			panic(common.ErrInvalidRequest(nil))
		}

		store := nodeStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := nodeBiz.NewGetResultBiz(store)
		result, err := biz.GetResult(c.Request.Context(), data)
		if err != nil {
			c.JSON(400, err)
			return
		}

		c.JSON(200, common.SimpleSuccessResponse(result))
	}
}
