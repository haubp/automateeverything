package commandTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/command/commandBiz"
	"backend_autotest/modules/command/commandModel"
	"backend_autotest/modules/command/commandStorage"
	"github.com/gin-gonic/gin"
)

func NewNodeCommand(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data commandModel.CommandNode
		data.User = c.MustGet(component.CurrentUser).(component.Requester).GetUserName()
		
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := commandStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := commandBiz.NewNodeCommandBiz(store)

		err := biz.AddNodeCommand(c.Request.Context(), &data)
		if err != nil {
			c.JSON(400, err)
		}

		c.JSON(200, common.SimpleSuccessResponse("add data success"))
	}
}
