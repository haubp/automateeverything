package userTransport

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"fmt"
	"github.com/gin-gonic/gin"
)

func PostTemplate(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		dataID := c.MustGet(component.CurrentUser).(component.Requester).GetUserName()

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.SaveUploadedFile(fileHeader, fmt.Sprintf("./fileRepository/users/%s", dataID))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(200, common.SimpleSuccessResponse("success"))
	}
}
