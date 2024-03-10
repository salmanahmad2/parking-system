package main

import (
	"parking/server/api"

	"github.com/gin-gonic/gin"
)

const defaultPort = ":8000"

func main() {
	router := gin.Default()

	api.Init()

	api.Routes(router)
	router.Run(defaultPort)
}
