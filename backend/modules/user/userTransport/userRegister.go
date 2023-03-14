package userTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/user/userBiz"
	"backend_autotest/modules/user/userModel"
	"backend_autotest/modules/user/userStorage"
	"github.com/gin-gonic/gin"
)

func UserRegister(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data userModel.User
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		//
		store := userStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := userBiz.NewCreateUserBiz(store)
		if err := biz.CreateNewUser(c.Request.Context(), &data); err != nil {
			c.JSON(400, err)
			return
		}
		c.JSON(200, common.SimpleSuccessResponse(data.UserName))

	}
}
