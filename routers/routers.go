package routers

import (
	"calculator_backend/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	router := r.Group("")
	{
		router.GET("/get_history", controllers.GetHistory)
		router.GET("/add_history", controllers.AddHistory)
	}
}
