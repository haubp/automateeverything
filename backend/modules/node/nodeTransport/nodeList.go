package nodeTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/node/nodeBiz"
	"backend_autotest/modules/node/nodeStorage"
	"backend_autotest/modules/user/userModel"

	"github.com/gin-gonic/gin"
)

func NodeList(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data userModel.User

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := nodeStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := nodeBiz.NewListNodeBiz(store)
		result, err := biz.ListNode(c.Request.Context(), &data)
		if err != nil {
			c.JSON(400, err)
			return
		}

		c.JSON(200, common.SimpleSuccessResponse(result))
	}
}
