package main

import (
	"fmt"
	"main/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gin.DisableConsoleColor()
	router.Static("/images", "./uploaded/images")
	api.Setup(router)

	// In case of running server on Heroku
	// router.Run(":8081") -> run in local
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("No Port In heroku")
	} else {
		fmt.Println("Environment port : " + port)
		router.Run(port)
	}
}
