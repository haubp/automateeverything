package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"

	"backend_autotest/common"
	"backend_autotest/component"
	"backend_autotest/middleware"
	"backend_autotest/modules/command/commandTransport"
	"backend_autotest/modules/node/nodeTransport"
	"backend_autotest/modules/user/userTransport"
)

func main() {

	db := common.InitMongoDB()

	redis := common.InitRedis()

	fmt.Println(db)

	if err := runService(db, redis); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *mongo.Client, redis *redis.Client) error {
	r := gin.Default()
	appCtx := component.NewAppContext(db, "abc&1*~#^2^#s0^=)^^7%b34", redis)

	user := r.Group("/user")
	{
		user.POST("/register", userTransport.UserRegister(appCtx))
		user.POST("/login", userTransport.UserLogin(appCtx))
		user.GET("/profile", middleware.RequireAuth(appCtx), userTransport.GetProfile(appCtx))

		user.GET("/log/agent", userTransport.GetAgentLogFile(appCtx))
		user.GET("/log/auto", userTransport.GetAutoLogFile(appCtx))

		user.POST("/template", middleware.RequireAuth(appCtx), userTransport.PostTemplate(appCtx))

	}

	node := r.Group("/node")
	{
		node.POST("/register", nodeTransport.NodeRegister(appCtx))
		node.DELETE("/delete", nodeTransport.NodeDelete(appCtx))
		node.GET("/list", nodeTransport.NodeList(appCtx))
		node.POST("/result", nodeTransport.NodePostResult(appCtx))
		node.GET("/result", nodeTransport.NodeGetResult(appCtx))

		node.POST("/log/agent", nodeTransport.PostAgentLogFile(appCtx))
		node.POST("/log/auto", nodeTransport.PostAutoLogFile(appCtx))
		node.GET("/template", nodeTransport.GetTemplate(appCtx))
	}

	command := node.Group("/command")
	{
		command.POST("/new", middleware.RequireAuth(appCtx), commandTransport.NewNodeCommand(appCtx))
		command.GET("/get", commandTransport.GetAndDeleteCommand(appCtx))

	}

	return r.Run(":8080")
}
