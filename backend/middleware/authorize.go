package middleware

import (
	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/component/generatetoken"
	"backend_autotest/modules/user/userModel"
	"backend_autotest/modules/user/userStorage"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	// split Bearer and token
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

func RequireAuth(appCtx component.AppContext) func(c *gin.Context) {

	tokenProvider := generatetoken.NewProviderToken(generatetoken.Payload{}, appCtx.GetSecret())

	return func(c *gin.Context) {

		redis := appCtx.GetRedis()

		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		val, err := redis.Get(token).Result()

		if err != nil {
			c.JSON(401, common.ErrInvalidLogin(err))
		} else {
			var user userModel.User
			user.UserName = val
			c.Set(component.CurrentUser, &user)
			c.Next()
		}

		db := appCtx.GetNewDataMongoDB()
		store := userStorage.NewMongoStore(db)

		if err := tokenProvider.Decrypt(token); err != nil {
			panic(err)
		}

		tmp, err := store.FindUser(c.Request.Context(), bson.M{"user_name": tokenProvider.Core.UserName})
		if err != nil {
			panic(err)
		}

		var user generatetoken.Payload
		user.UserName = tmp.UserName

		c.Set(component.CurrentUser, user)
		c.Next()
	}
}
