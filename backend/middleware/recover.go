package middleware

import (
	"github.com/gin-gonic/gin"

	"backend_autotest/common"
	"backend_autotest/component"
)

func Recover(ac component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// if catch error, custom Header
				c.Header("Content-Type", "application/json")

				// case 1: catch error, catch AppError
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// catch panic of gin
					panic(err)
					return
				}

				// case 2: if catch error not AppError, custom err end return
				// this error is not in appError, so we can't handle it
				appErr := common.ErrInternal(err.(error))
				// AbortWithStatusJson is destructor function, We must stop
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				// catch panic of gin
				panic(err)
				return
			}
		}()

		// Call the next handler, if exist next middleware, it will be go to next middleware, if not, it will go to handler
		// this is endpoint
		c.Next()
	}
}
