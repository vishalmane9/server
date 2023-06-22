package main

import (
	"flag"

	"github.com/gin-gonic/gin"
)

func main() {
	option := flag.String("admin", "restserver", "communication between rest rest client and server")
	flag.Parse()

	switch *option {
	case "restclient":
		runRestClient()
	case "restserver":
		router := gin.Default()
		router.POST("/vishal", runRestServer)
		err := router.Run(":8082")
		if err != nil {
			panic("[Error] failed to start Gin server due to: " + err.Error())
		}
	case "grpcserver":
		runGrpcServer()
	}
}
