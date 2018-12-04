package routes

import (
	"gin2018/app/http/controllers"

	"github.com/gin-gonic/gin"
)

type Api struct{}

func (Api) Include(router *gin.Engine) {
	router.GET("/api/hotel", controllers.HotelController{}.Index)
	router.POST("/api/hotel", controllers.HotelController{}.Store)
	router.GET("/api/hotel/:id", controllers.HotelController{}.Show)
	router.PUT("/api/hotel/:id", controllers.HotelController{}.Update)
	router.DELETE("/api/hotel/:id", controllers.HotelController{}.Destroy)
}
