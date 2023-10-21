package main

import (
	"calculator_backend/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	routers.InitRouter(router)

	err := router.Run(":2333")
	if err != nil {
		log.Printf("Having problem starting gin.Engine: %v \n", err)
		return
	}
}
