package userTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(appCtx component.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(component.CurrentUser).(component.Requester)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
