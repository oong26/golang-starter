package startup

import (
	reportgenerator "my-golang-starter/report-generator/api/example"
	"my-golang-starter/report-generator/config"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	env := config.NewEnv(".env", true)
	router := gin.Default()
	CreateRoutes(router)
	var host string = env.ServerHost + ":" + env.ServerPort
	router.Run(host)
}

func CreateRoutes(router *gin.Engine) {
	router.GET("/", reportgenerator.Welcome)
}
