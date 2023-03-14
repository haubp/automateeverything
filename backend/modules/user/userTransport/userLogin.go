package userTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/modules/user/userBiz"
	"backend_autotest/modules/user/userModel"
	"backend_autotest/modules/user/userStorage"
	"github.com/gin-gonic/gin"
	"time"
)

func UserLogin(app component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data userModel.UserLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userStorage.NewMongoStore(app.GetNewDataMongoDB())
		biz := userBiz.NewLoginUserBiz(store)
		token, err := biz.LoginUser(c.Request.Context(), &data, app.GetSecret())
		if err != nil {
			c.JSON(400, err)
			return
		}

		redis := app.GetRedis()

		if err := redis.Set(token, data.UserName, 24*time.Hour).Err(); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(token))
	}
}
